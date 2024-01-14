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
	"github.com/renatormc/pfila/api/database/repo"
	"github.com/renatormc/pfila/api/processes"
)

type Runner struct {
	Cmd     *exec.Cmd
	Outfile *os.File
	ProcID  int64
}

func NewRunner(procID int64) *Runner {
	return &Runner{ProcID: procID}
}

func (r *Runner) Cancel() {
	if err := r.Cmd.Process.Kill(); err != nil {
		log.Println(err)
		return
	}
	r.RegisterFinish("ERRO", "Processo cancelado pelo usuário.")
}

func (r *Runner) RegisterFinish(status string, message string) {
	proc := repo.GetProcessByIdOrFail(r.ProcID)
	proc.Status = status
	proc.Finish = time.Now()
	if err := repo.SaveProc(proc); err != nil {
		log.Fatal(err)
	}
	if message != "" {
		r.Outfile.WriteString(message)
	}
}

func (r *Runner) Run() {
	log.Printf("start running %d\n", r.ProcID)
	cf := config.GetConfig()
	proc := repo.GetProcessByIdOrFail(r.ProcID)
	var err error
	r.Outfile, err = os.Create(filepath.Join(cf.ConsoleFolder, proc.RandomID))
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		r.Outfile.Close()
		os.Exit(0)
	}()
	args, err := processes.GetCmdArgs(proc)
	if err != nil {
		r.Outfile.WriteString(err.Error())
		log.Println(err)
		r.RegisterFinish("ERRO", err.Error())
		return
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = r.Outfile
	cmd.Stdout = r.Outfile
	cmd.Stdin = os.Stdin

	proc.Status = "EXECUTANDO"
	proc.Start = time.Now()
	err = repo.SaveProc(proc)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Run(); err != nil {
		log.Println(err)
		r.Outfile.WriteString("\nPFila: Finalizou com erro")
		proc, err := repo.GetProcessById(int64(proc.ID))
		if err != nil {
			r.Outfile.WriteString(fmt.Sprintf("Process of id %d not found", proc.ID))
			log.Printf("Process of id %d not found", proc.ID)
			return
		}
		r.RegisterFinish("ERRO", "Erro ao iniciar ao processo")
		return
	}

	r.RegisterFinish("FINALIZADO", "Processo finalizado.")
}

func Serve(port int, runner *Runner) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"procID": runner.ProcID})
	})

	r.GET("/cancel", func(c *gin.Context) {
		go func() {
			time.Sleep(5 * time.Second)
			runner.Cancel()
		}()
		c.JSON(http.StatusOK, gin.H{"procID": runner.ProcID})
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
	runner := NewRunner(int64(*procID))
	go Serve(*port, runner)
	runner.Run()
}
