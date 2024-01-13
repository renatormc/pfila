package processes

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"strings"

	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/database/repo"
	"github.com/renatormc/pfila/api/helpers"

	"github.com/renatormc/pfila/api/utils"

	"gorm.io/gorm"
)

func IsProcessRunning(proc *models.Process) bool {
	log.Fatalln("Not implemented")
	return true
}

func Run(proc *models.Process) {
	cf := config.GetConfig()
	proc.Port = helpers.GetFreePort()
	cmd := exec.Command(filepath.Join(cf.AppDir, "pfila_runner"), "--port", fmt.Sprintf("%d", proc.Port), "--proc-id", fmt.Sprintf("%d", proc.ID))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		proc.Status = "ERRO"
		proc.Finish = time.Now()
		log.Println(err)
	}
	if err := cmd.Process.Release(); err != nil {
		log.Println(err)
	}
	db := database.GetDatabase()
	if err := db.Save(proc).Error; err != nil {
		log.Fatal(err)
	}
}

func StopProcess(proc *models.Process) error {
	log.Fatalln("not implemented")
	return nil
	// cmd := GetRunningCmds().GetRunningCmd(proc.ID)
	// if cmd == nil {
	// 	log.Fatalf("not found process %d", proc.ID)
	// }

	// if err := cmd.Process.Kill(); err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// GetRunningCmds().DeleteRunningCmd(proc.ID)

	// if proc.IsDocker {
	// 	err := exec.Command("docker", "stop", proc.RandomID).Run()
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// return nil
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
			if p.Status != "FINALIZADO" {
				allFinished = false
				break
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

func HasDependencyNotFinished(proc *models.Process) (bool, error) {
	db := database.GetDatabase()
	deps := proc.GetDependencies()
	var count int64
	if err := db.Model(&models.Process{}).Where("status != ?", "FINALIZADO").Where("id IN ?", deps).Count(&count).Error; err != nil {
		log.Println(err)
		return false, err
	}
	return count > 0, nil
}

func GetNext() (*models.Process, error) {
	db := database.GetDatabase()
	proc := models.Process{}
	err := db.Where("status = ?", "AGUARDANDO").Order("start_waiting asc").First(&proc).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &proc, nil
}

func CheckRunningProcesses() bool {
	db := database.GetDatabase()
	procs := []models.Process{}
	err := db.Where("status = ?", "EXECUTANDO").Find(&procs).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Fatal(err)
	}
	ret := false
	if len(procs) > 0 {
		for _, proc := range procs {
			if IsProcessRunning(&proc) {
				ret = true
			} else {
				proc.Status = "ERRO"
				proc.Finish = time.Now()
				if err := db.Save(&proc).Error; err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	return ret
}

func CheckProcesses() error {
	if CheckRunningProcesses() {
		return nil
	}

	if err := CheckWaitingDep(); err != nil {
		return err
	}

	for {
		proc, err := GetNext()
		if err != nil {
			return err
		}
		if proc == nil {
			break
		}
		hasDep, err := HasDependencyNotFinished(proc)
		if err != nil {
			return err
		}
		if hasDep {
			proc.Status = "AGUARDANDO_DEP"
			if err := repo.SaveProc(proc); err != nil {
				return err
			}
		} else {
			Run(proc)
			break
		}
	}
	return nil
}
