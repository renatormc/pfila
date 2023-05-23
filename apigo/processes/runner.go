package processes

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/database/repo"
)

func Runner(proc models.Process) {
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
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = outfile
	cmd.Stdout = outfile
	cmd.Stdin = os.Stdin
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start command: %v\n", err)
	}
	if err := cmd.Wait(); err != nil {
		outfile.WriteString("\nPFila: Finalizou com erro")
		procNew, err := repo.GetProcessById(int64(proc.ID))
		if err != nil {
			outfile.WriteString(fmt.Sprintf("Process of id %d not found", proc.ID))
			log.Fatalf("Process of id %d not found", proc.ID)
		}
		procNew.Status = "ERRO"
		procNew.Finish = time.Now()
		repo.SaveProc(procNew)
		log.Println(err)
	}

	procNew, err := repo.GetProcessById(int64(proc.ID))
	if err != nil {
		outfile.WriteString(fmt.Sprintf("Process of id %d not found", proc.ID))
		log.Fatalf("Process of id %d not found", proc.ID)
	}
	proc.Status = "FINALIZADO"
	proc.Finish = time.Now()
	if err := repo.SaveProc(procNew); err != nil {
		log.Println(err)
	}
}
