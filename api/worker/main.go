package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/database/repo"
	"github.com/renatormc/pfila/api/processes"
	"gorm.io/gorm"
)

func RunNext() {
	db := database.GetDatabase()
	proc := &models.Process{}
	err := db.Where("status = ?", "AGUARDANDO").Order("start_waiting asc").First(proc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		log.Fatal(err)
	}
	cf := config.GetConfig()
	outfile, err := os.Create(filepath.Join(cf.ConsoleFolder, proc.RandomID))
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	args, err := processes.GetCmdArgs(proc)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = outfile
	cmd.Stdout = outfile
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
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

func main() {
	for {
		RunNext()
		time.Sleep(5 * time.Second)
	}
}
