package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/akamensky/argparse"
	"github.com/gin-gonic/gin"
	"github.com/renatormc/pfila/api/database/repo"
	"github.com/renatormc/pfila/api/processes"
)

type Runner struct {
	Cmd    *exec.Cmd
	ProcID int64
}

func NewRunner(procID int64) *Runner {
	return &Runner{ProcID: procID}
}

func (r *Runner) Cancel() {
	r.RegisterFinish("ERRO", "Processo cancelado pelo usuário.")
	if err := r.Cmd.Process.Kill(); err != nil {
		log.Println(err)
		return
	}

}

func (r *Runner) RegisterFinish(status string, message string) {
	proc := repo.GetProcessByIdOrFail(r.ProcID)
	proc.Status = status
	proc.Finish = time.Now()
	if err := repo.SaveProc(proc); err != nil {
		log.Fatal(err)
	}
	if message != "" {
		fmt.Println(message)
	}
}

func (r *Runner) Run() {
	log.Printf("start running %d\n", r.ProcID)
	proc := repo.GetProcessByIdOrFail(r.ProcID)
	var err error
	defer func() {
		os.Exit(0)
	}()
	args, err := processes.GetCmdArgs(proc)
	if err != nil {
		log.Println(err)
		r.RegisterFinish("ERRO", err.Error())
		return
	}
	r.Cmd = exec.Command(args[0], args[1:]...)
	r.Cmd.Stderr = os.Stderr
	r.Cmd.Stdout = os.Stdout

	if err := r.Cmd.Start(); err != nil {
		log.Println(err)
		proc, err := repo.GetProcessById(int64(proc.ID))
		if err != nil {
			log.Printf("Process of id %d not found", proc.ID)
			return
		}
		return
	}

	proc.Status = "EXECUTANDO"
	proc.Start = time.Now()
	proc.Pid = r.Cmd.Process.Pid
	err = repo.SaveProc(proc)
	if err != nil {
		log.Fatal(err)
	}

	if err := r.Cmd.Wait(); err != nil {
		log.Println(err)
		proc, err := repo.GetProcessById(int64(proc.ID))
		if err != nil {
			log.Printf("Process of id %d not found", proc.ID)
			return
		}
		return
	}

	r.RegisterFinish("FINALIZADO", "Processo finalizado.")
}

func Serve(port int, runner *Runner) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"procID": runner.ProcID})
	})

	r.GET("/cancel/:procID", func(c *gin.Context) {
		procID, err := strconv.ParseInt(c.Param("procID"), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "wrong value"})
			return
		}
		if procID != runner.ProcID {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}
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
