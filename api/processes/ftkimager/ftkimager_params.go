package ftkimager

import (
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/renatormc/pfila/api/helpers"
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

func (p *FtkimagerParams) Validate(ve *helpers.ValidationError) {
	disks, err := GetDisks()
	if err != nil {
		log.Println(err)
		ve.AddMessage("internal", "internal error")
		return
	}
	if !helpers.SliceContains(disks, p.Disk) {
		ve.AddMessage("disk", "Disco não encontrado")
	}
	dest := strings.TrimSpace(p.Destination)
	if dest == "" {
		ve.AddMessage("destination", "Campo obrigatório")
	} else {
		folder := filepath.Dir(dest)
		if !filepath.IsAbs(folder) || !helpers.DirectoryExists(folder) {
			ve.AddMessage("destination", "Caminho inválido")
		}
	}

}
