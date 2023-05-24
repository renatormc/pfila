package database

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func CreateLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	return newLogger
}

func Connect() (*gorm.DB, error) {
	var err error
	cf := config.GetConfig()
	path := filepath.Join(cf.AppDir, "pfila.db")
	c, err := gorm.Open(sqlite.Open(path), &gorm.Config{Logger: CreateLogger()})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	conf, _ := c.DB()
	conf.SetMaxOpenConns(1)
	return c, nil
}

func CloseConn() error {
	if db != nil {
		config, err := db.DB()
		if err != nil {
			return err
		}

		err = config.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func GetDatabase() *gorm.DB {
	if db == nil {
		aux, err := Connect()
		if err != nil {
			log.Fatal(err)
		}
		db = aux
	}
	return db
}

func Migrate() {
	db := GetDatabase()
	db.AutoMigrate(&models.Process{})
}
