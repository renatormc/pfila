package processes

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/reantormc/pfila/api/config"
	"github.com/reantormc/pfila/api/database"
	"github.com/reantormc/pfila/api/database/models"
)

func Run(proc models.Process) error {
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
