package processes

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/database/repo"
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
	proc.Status = "RUNNING"
	proc.RandomID = uuid.NewString()

	db := database.GetDatabase()
	err = db.Save(&proc).Error
	if err != nil {
		return err
	}

	return nil
}

func StopProcess(proc *models.Process) error {
	p, err := os.FindProcess(proc.Pid)
	if err != nil {
		return nil
	}

	var rusage syscall.Rusage
	if err := syscall.Getrusage(proc.Pid, &rusage); err != nil {
		log.Println(err)
		return err
	}
	startTime := time.Unix(int64(rusage.Utime.Sec), int64(rusage.Utime.Usec)*1000)
	if math.Abs(startTime.Sub(proc.Start).Seconds()) < 30 {
		if err := p.Kill(); err != nil {
			log.Println(err)
			return err
		}
	}
	proc.Status = "CANCELED"
	proc.Finish = time.Now()
	if err := repo.SaveProc(proc); err != nil {
		log.Println(err)
		return err
	}
	return nil
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

	return strings.Join(lines, "\n")
}

func CheckProcesses() error {
	var n int64
	db := database.GetDatabase()
	if err := db.Model(models.Process{}).Where("status = ?", "RUNNING").Count(&n).Error; err != nil {
		return err
	}
	if n == 0 {
		proc := models.Process{}
		err := db.Where("status = ?", "WAITING").Order("start_waiting asc").First(&proc).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
		return Run(&proc)
	}
	return nil
}
