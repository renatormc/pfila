package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/akamensky/argparse"
	"github.com/gin-gonic/gin"
	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/database/repo"
	"github.com/renatormc/pfila/api/processes"
)

func Run(proc *models.Process) {
	log.Printf("start running %d\n", proc.ID)
	cf := config.GetConfig()
	outfile, err := os.Create(filepath.Join(cf.ConsoleFolder, proc.RandomID))
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		outfile.Close()
		os.Exit(0)
	}()
	args, err := processes.GetCmdArgs(proc)
	if err != nil {
		outfile.WriteString(err.Error())
		log.Println(err)
		proc.Status = "ERRO"
		proc.Finish = time.Now()
		err = repo.SaveProc(proc)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = outfile
	cmd.Stdout = outfile
	cmd.Stdin = os.Stdin

	proc.Status = "EXECUTANDO"
	proc.Start = time.Now()
	err = repo.SaveProc(proc)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Run(); err != nil {
		log.Println(err)
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

func Serve(port int, procID int) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"procID": procID})
	})

	log.Fatal(r.Run(fmt.Sprintf(":%d", port)))
}

func main() {
	parser := argparse.NewParser("pfila_runner", "Runs a process")
	port := parser.Int("p", "port", &argparse.Options{Required: true, Help: "Port to run application"})
	procID := parser.Int("i", "proc-id", &argparse.Options{Required: true, Help: "Process ID"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	Serve(*port, *procID)
}
