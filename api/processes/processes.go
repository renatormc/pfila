package processes

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"strings"
	"time"

	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/database/repo"

	"github.com/renatormc/pfila/api/utils"

	"gorm.io/gorm"
)

func WriteErrorToConsole(err error, proc *models.Process) error {
	log.Println(err)
	cf := config.GetConfig()
	outfile := filepath.Join(cf.ConsoleFolder, proc.RandomID)
	text := fmt.Sprintf("%s\n#pfilaerror#", err.Error())
	if err := os.WriteFile(outfile, []byte(text), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	db := database.GetDatabase()
	proc.Status = "ERRO"
	proc.Finish = time.Now()
	err = db.Save(&proc).Error
	if err != nil {
		return err
	}
	return nil
}

func Run(proc *models.Process) {
	log.Printf("start running %d\n", proc.ID)
	cf := config.GetConfig()
	outfile, err := os.Create(filepath.Join(cf.ConsoleFolder, proc.RandomID))
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		outfile.Close()
		DeleteRunningCmd(proc.ID)
	}()
	args, err := GetCmdArgs(proc)
	if err != nil {
		outfile.WriteString(err.Error())
		log.Println(err)
		proc.Status = "ERRO"
		proc.Finish = time.Now()
		err = repo.SaveProc(proc)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	cmd := exec.Command(args[0], args[1:]...)
	SaveRunningCmd(proc.ID, cmd)
	cmd.Stderr = outfile
	cmd.Stdout = outfile
	cmd.Stdin = os.Stdin

	proc.Status = "EXECUTANDO"
	proc.Start = time.Now()
	err = repo.SaveProc(proc)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Run(); err != nil {
		log.Println(err)
		outfile.WriteString("\nPFila: Finalizou com erro")
		proc, err := repo.GetProcessById(int64(proc.ID))
		if err != nil {
			outfile.WriteString(fmt.Sprintf("Process of id %d not found", proc.ID))
			log.Printf("Process of id %d not found", proc.ID)
			return
		}
		proc.Status = "ERRO"
		proc.Finish = time.Now()
		err = repo.SaveProc(proc)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	proc, err = repo.GetProcessById(int64(proc.ID))
	if err != nil {
		outfile.WriteString(fmt.Sprintf("Process of id %d not found", proc.ID))
		log.Printf("Process of id %d not found", proc.ID)
		return
	}
	proc.Status = "FINALIZADO"
	proc.Finish = time.Now()
	err = repo.SaveProc(proc)
	if err != nil {
		log.Fatal(err)
	}
}

func StopProcess(proc *models.Process) error {
	cmd := GetRunningCmd(proc.ID)
	if cmd == nil {
		log.Fatalf("not found process %d", proc.ID)
	}

	if err := cmd.Process.Kill(); err != nil {
		log.Println(err)
		return err
	}
	DeleteRunningCmd(proc.ID)

	if proc.IsDocker {
		err := exec.Command("docker", "stop", proc.RandomID).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

type Status string

const (
	StatusError   Status = "error"
	StatusOk      Status = "ok"
	StatusUnknown Status = "unknown"
)

func CheckConsoleEndMessage(proc *models.Process) Status {
	cf := config.GetConfig()
	path := filepath.Join(cf.ConsoleFolder, proc.RandomID)
	text, err := utils.ReadTail(path, 20)
	if err == nil {
		if strings.Contains(text, "#pfilaok#") {
			return StatusOk
		}
		if strings.Contains(text, "#pfilaerror#") {
			return StatusError
		}
	}
	return StatusUnknown
}

func GetProcConsole(proc *models.Process, nLines int) string {
	cf := config.GetConfig()
	path := filepath.Join(cf.ConsoleFolder, proc.RandomID)
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	buf := make([]string, nLines)
	for scanner.Scan() {
		copy(buf, buf[1:])
		buf[len(buf)-1] = scanner.Text()
	}

	if len(buf) > nLines {
		lines = buf[len(buf)-nLines:]
	} else {
		lines = buf
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

// func CheckRunningProcs() error {
// 	db := database.GetDatabase()
// 	procs := []models.Process{}
// 	if err := db.Where("status = ?", "EXECUTANDO").Find(&procs).Error; err != nil {
// 		if !errors.Is(err, gorm.ErrRecordNotFound) {
// 			return err
// 		}
// 	}
// 	for _, proc := range procs {
// 		_, err := helpers.GetProcess(int32(proc.Pid), proc.Start)
// 		if err != nil {
// 			res := CheckConsoleEndMessage(&proc)
// 			if res == StatusOk {
// 				proc.Status = "FINALIZADO"
// 			} else {
// 				proc.Status = "ERRO"
// 			}
// 			proc.Finish = time.Now()
// 			err = repo.SaveProc(&proc)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

func CheckWaitingDep() error {
	db := database.GetDatabase()
	procs := []models.Process{}
	if err := db.Where("status = ?", "AGUARDANDO_DEP").Find(&procs).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	deps := []models.Process{}
	for _, proc := range procs {
		if err := db.Find(&deps, proc.GetDependencies()).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
		}
		allFinished := true
		for _, p := range deps {
			if p.Status == "ERRO" {
				proc.Status = "ERRO"
				proc.Finish = time.Now().Local()
				if err := db.Save(&proc).Error; err != nil {
					return err
				}
				break
			} else if p.Status != "FINALIZADO" {
				allFinished = false
			}
		}
		if allFinished {
			proc.Status = "AGUARDANDO"
			if err := db.Save(&proc).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func CheckProcesses() error {
	if err := CheckWaitingDep(); err != nil {
		return err
	}

	db := database.GetDatabase()
	procs := []models.Process{}
	if err := db.Where("status = ?", "EXECUTANDO").Find(&procs).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if len(procs) == 0 {
		proc := models.Process{}

		err := db.Where("status = ?", "AGUARDANDO").Order("start_waiting asc").First(&proc).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
		go Run(&proc)
	}
	return nil
}
