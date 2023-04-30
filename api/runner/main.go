package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/akamensky/argparse"
	"github.com/reantormc/pfila/api/config"
	"github.com/reantormc/pfila/api/database/repo"
)

func main() {

	parser := argparse.NewParser("PFila Runner", "PFila Runner")
	id := parser.Int("p", "id", &argparse.Options{Required: true, Help: "ID of the process to run"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	proc := repo.GetProcessById(int64(*id))
	if proc == nil {
		fmt.Printf("Process of id %d not found", *id)
		os.Exit(1)
	}

	cf := config.GetConfig()
	outfile, err := os.Create(filepath.Join(cf.ConsoleFolder, proc.RandomID))
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	cmd := exec.Command(filepath.Join(cf.ConsoleFolder, proc.RandomID))
	cmd.Stderr = outfile
	cmd.Stdout = outfile
	if err := cmd.Run(); err != nil {
		outfile.WriteString("\nPFila: Finalizou com erro")
		proc = repo.GetProcessById(int64(*id))
		if proc == nil {
			outfile.WriteString(fmt.Sprintf("Process of id %d not found", *id))
			log.Fatalf("Process of id %d not found", *id)
		}
		proc.Status = "FINISHED"
		proc.Finish = time.Now()
		repo.SaveProc(proc)
		log.Fatal(err)
	}

}
