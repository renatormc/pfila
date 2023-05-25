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

type CurrentProcess struct {
	cmd  *exec.Cmd
	proc *models.Process
}

type Worker struct {
	current *CurrentProcess
}

func (w *Worker) Loop() {
	for {
		if w.current == nil {
			w.PrepareNext()
			go w.RunNext()
		}
		time.Sleep(5 * time.Second)
	}
}

func (w *Worker) PrepareNext() {
	db := database.GetDatabase()
	proc := &models.Process{}
	err := db.Where("status = ?", "PROXIMO").Order("start_waiting asc").First(proc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		log.Fatal(err)
	}
	cf := config.GetConfig()
	outfilePath := filepath.Join(cf.ConsoleFolder, proc.RandomID)
	outfile, err := os.Create(outfilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	args, err := processes.GetCmdArgs(proc)
	if err != nil {
		log.Fatal(err)
	}
	w.cmd = exec.Command(args[0], args[1:]...)
}

func (w *Worker) RunNext() {
	cf := config.GetConfig()
	outfile, err := os.Create(filepath.Join(cf.ConsoleFolder, proc.RandomID))
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	args, err := GetCmdArgs(&proc)
	if err != nil {
		outfile.WriteString(err.Error())
		log.Println(err)
	}
	w.cmd.Stderr = os.Stderr
	w.cmd.Stdout = os.Stdout
	w.cmd.Stdin = os.Stdin
	log.Printf("Starting %d\n", proc.ID)
	if err := cmd.Start(); err != nil {
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

	proc.Pid = cmd.Process.Pid
	proc.Status = "EXECUTANDO"
	proc.Start = time.Now()
	err = repo.SaveProc(proc)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	w := Worker{}
	w.Loop()
}
