package iped

import (
	"os/exec"
	"path/filepath"

	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/helpers"
)

type IpedParams struct {
	Destination string   `json:"destination"`
	Sources     []string `json:"sources"`
	Portable    bool     `json:"portable"`
	Profile     string   `json:"profile"`
}

func (p *IpedParams) ToCmd() *exec.Cmd {
	args := []string{}
	for _, src := range p.Sources {
		args = append(args, "-d")
		args = append(args, src)
	}
	args = append(args, "-o")
	args = append(args, p.Destination)
	if p.Portable {
		args = append(args, "--portable")
	}
	args = append(args, "-profile")
	args = append(args, p.Profile)
	return exec.Command("iped", args...)
}

func (p *IpedParams) Validate(ve *helpers.ValidationError) {
	if !helpers.DirectoryExists(p.Destination) {
		ve.AddMessage("destination", "Diretório não encontrado")
	}
	for _, src := range p.Sources {
		if !helpers.DirectoryExists(src) || !helpers.FileExists(src) {
			ve.AddMessage("sources", "Fonte não encontrada")
			break
		}
	}
	cf := config.GetConfig()
	path := filepath.Join(cf.IpedProfileFolder, p.Profile)
	if !helpers.DirectoryExists(path) {
		ve.AddMessage("profile", "Perfil não encontrado")
	}
}
