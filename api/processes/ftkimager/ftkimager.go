package ftkimager

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func GetDisks() ([]string, error) {
	disks := []string{}
	cmd := exec.Command("ftkimager", "--list-drives")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return disks, err
	}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		text := strings.TrimSpace(line)
		if runtime.GOOS == "windows" {
			if strings.HasPrefix(text, "\\\\.\\PHYSICALDRIVE") {
				disks = append(disks, text)
			}
		} else {
			if strings.HasPrefix(text, "/dev") {
				disks = append(disks, text)
			}
		}

	}
	fmt.Println(string(output))
	return disks, nil
}