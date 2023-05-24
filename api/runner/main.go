package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	outfile, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stderr = outfile
	cmd.Stdout = outfile
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
