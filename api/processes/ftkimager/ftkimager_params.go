package ftkimager

import (
	"os/exec"
)

type FtkimagerParams struct {
	Disk        string `json:"disk"`
	Destination string `json:"destination"`
	Verify      bool   `json:"verify"`
	Format      string `json:"format"`
}

func (p *FtkimagerParams) ToCmd() *exec.Cmd {
	args := []string{p.Disk, p.Destination}
	if p.Format == "e01" {
		args = append(args, "--e01")
	}
	if p.Verify {
		args = append(args, "--verify")
	}
	return exec.Command("ftkimager", args...)
}
