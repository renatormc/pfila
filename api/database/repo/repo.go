package repo

import (
	"errors"
	"log"

	"github.com/renatormc/pfila/api/database"
	"github.com/renatormc/pfila/api/database/models"
	"gorm.io/gorm"
)

func GetProcessById(id int64) (*models.Process, error) {
	db := database.GetDatabase()
	proc := models.Process{}
	if err := db.First(&proc, id).Error; err != nil {
		return nil, err
	}
	return &proc, nil
}

func GetAllProcesses() []models.Process {
	db := database.GetDatabase()
	procs := []models.Process{}
	if err := db.Order("created_at asc").Find(&procs).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal(err)
		}

	}
	return procs
}

func SaveProc(proc *models.Process) error {
	db := database.GetDatabase()
	if err := db.Save(&proc).Error; err != nil {
		return err
	}
	return nil
}

func GetProcessesByStatus(status string) []models.Process {
	db := database.GetDatabase()
	procs := []models.Process{}
	if err := db.Where("status = ?", status).Find(&procs).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal(err)
		}
	}
	return procs
}
