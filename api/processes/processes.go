package processes

import (
	"bufio"
	"errors"
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

	"github.com/renatormc/pfila/api/utils"

	"github.com/shirou/gopsutil/process"
	"gorm.io/gorm"
)

func GetProcessStartTime(pid int32) (time.Time, error) {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return time.Time{}, err
	}
	proc.Kill()
	createTime, err := proc.CreateTime()
	if err != nil {
		return time.Time{}, err
	}

	startTime := time.Unix(createTime, 0)
	return startTime, nil
}

type Resp struct {
	ProcID int `json:"procID"`
}

func StopProcess(proc *models.Process) error {
	p, err := process.NewProcess(int32(proc.Pid))
	if err != nil {
		return err
	}
	if err := p.Kill(); err != nil {
		return err
	}
	return nil
}

func IsProcessRunning(proc *models.Process) bool {
	t, err := GetProcessStartTime(int32(proc.Pid))
	if err != nil {
		log.Fatal(err)
	}
	delta := proc.Start.Sub(t)
	threshold := time.Second * 2
	if delta > -threshold && delta < threshold {
		return true
	}
	return false
}

func Run(proc *models.Process) {
	cf := config.GetConfig()
	outputFile := filepath.Join(cf.ConsoleFolder, proc.RandomID)
	scriptExe := filepath.Join(cf.AppDir, "pfila_runner.sh")
	cmd := exec.Command(scriptExe, outputFile)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		proc.Status = "ERRO"
		proc.Finish = time.Now()
		log.Println(err)
	}
	db := database.GetDatabase()

	if err := db.Save(proc).Error; err != nil {
		log.Fatal(err)
	}
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

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
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
