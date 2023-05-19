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

	"github.com/google/uuid"
	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/database/repo"
	"github.com/renatormc/pfila/api/helpers"
	"gorm.io/gorm"
)

func Run(proc *models.Process) error {
	cf := config.GetConfig()
	cmd := exec.Command(filepath.Join(cf.AppDir, "pfila_runner"), "-p", fmt.Sprintf("%d", proc.ID))
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Process.Release()
	if err != nil {
		return err
	}
	proc.Pid = cmd.Process.Pid
	proc.Start = time.Now()
	proc.Status = "EXECUTANDO"
	proc.RandomID = uuid.NewString()

	db := database.GetDatabase()
	err = db.Save(&proc).Error
	if err != nil {
		return err
	}

	return nil
}

func StopProcess(proc *models.Process) error {
	p, err := helpers.GetProcess(int32(proc.Pid), proc.Start)
	if err != nil {
		return nil
	}

	if err := p.Kill(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetProcConsole(proc *models.Process, nLines int) string {
	cf := config.GetConfig()
	path := filepath.Join(cf.ConsoleFolder, proc.RandomID)
	log.Println(path)
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

func CheckProcesses() error {

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
		return Run(&proc)
	} else {
		for _, proc := range procs {
			_, err := helpers.GetProcess(int32(proc.Pid), proc.Start)
			if err != nil {
				proc.Status = "ERRO"
				err = repo.SaveProc(&proc)
				if err != nil {
					return err
				}
			}
		}

	}
	return nil
}
