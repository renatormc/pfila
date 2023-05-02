package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/akamensky/argparse"
	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database/repo"
	"github.com/renatormc/pfila/api/processes"
)

func main() {

	parser := argparse.NewParser("PFila Runner", "PFila Runner")
	id := parser.Int("p", "id", &argparse.Options{Required: true, Help: "ID of the process to run"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	proc, err := repo.GetProcessById(int64(*id))
	if err != nil {
		fmt.Printf("Process of id %d not found", *id)
		os.Exit(1)
	}

	cf := config.GetConfig()
	outfile, err := os.Create(filepath.Join(cf.ConsoleFolder, proc.RandomID))
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	cmd, err := processes.GetCmd(proc)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stderr = outfile
	cmd.Stdout = outfile
	if err := cmd.Run(); err != nil {
		outfile.WriteString("\nPFila: Finalizou com erro")
		proc, err := repo.GetProcessById(int64(*id))
		if err != nil {
			outfile.WriteString(fmt.Sprintf("Process of id %d not found", *id))
			log.Fatalf("Process of id %d not found", *id)
		}
		proc.Status = "ERROR"
		proc.Finish = time.Now()
		repo.SaveProc(proc)
		log.Fatal(err)
	}
	proc, err = repo.GetProcessById(int64(*id))
	if err != nil {
		outfile.WriteString(fmt.Sprintf("Process of id %d not found", *id))
		log.Fatalf("Process of id %d not found", *id)
	}
	proc.Status = "FINISHED"
	proc.Finish = time.Now()
	repo.SaveProc(proc)

}
