package repo

import (
	"errors"
	"log"

	"github.com/reantormc/pfila/api/database"
	"github.com/reantormc/pfila/api/database/models"
	"gorm.io/gorm"
)

func GetProcessById(id int64) *models.Process {
	db := database.GetDatabase()
	proc := models.Process{}
	if err := db.First(&proc, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		log.Fatal(err)
	}
	return &proc
}

func SaveProc(proc *models.Process) {
	db := database.GetDatabase()
	if err := db.Save(&proc).Error; err != nil {
		log.Fatal(err)
	}
}
