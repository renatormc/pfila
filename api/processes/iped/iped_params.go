package iped

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/helpers"
	"github.com/renatormc/pfila/api/utils"
)

type IpedParams struct {
	Destination string   `json:"destination"`
	Sources     []string `json:"sources"`
	Portable    bool     `json:"portable"`
	Profile     string   `json:"profile"`
}

func (p *IpedParams) ToCmdWindows() (*exec.Cmd, error) {
	cf := config.GetConfig()
	java := filepath.Join(cf.IpedFolder, "jre", "bin", "java.exe")
	jar := filepath.Join(cf.IpedFolder, "iped.jar")
	args := []string{"-jar", jar}
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
	args = append(args, "--nogui")
	return exec.Command(java, args...), nil
}

func (p *IpedParams) ToCmdLinux() (*exec.Cmd, error) {
	cf := config.GetConfig()
	args := []string{"run", "--rm"}
	args = append(args, "-v")
	args = append(args, fmt.Sprintf("%s://opt/IPED/iped-4.1.1/profiles", cf.IpedFolder))
	for _, src := range p.Sources {
		args = append(args, "-v")
		args = append(args, fmt.Sprintf("%s:/evidences%s", src, src))
	}
	args = append(args, "-v")
	args = append(args, fmt.Sprintf("%s:/evidences%s", p.Destination, p.Destination))
	args = append(args, "ipeddocker/iped:processor_4.1.1_3")
	args = append(args, "java")
	args = append(args, "-jar")
	args = append(args, "iped.jar")
	args = append(args, "--nogui")
	for _, src := range p.Sources {
		args = append(args, "-d")
		args = append(args, fmt.Sprintf("/evidences%s", src))
	}
	args = append(args, "-o")
	args = append(args, fmt.Sprintf("/evidences%s", p.Destination))

	if p.Portable {
		args = append(args, "--portable")
	}
	args = append(args, "-profile")
	args = append(args, p.Profile)
	return exec.Command("docker", args...), nil
}

func (p *IpedParams) ToCmd() (*exec.Cmd, error) {
	if runtime.GOOS == "windows" {
		return p.ToCmdWindows()
	}
	return p.ToCmdLinux()
}

func (p *IpedParams) Validate(ve *helpers.ValidationError) {
	if !utils.DirectoryExists(p.Destination) {
		ve.AddMessage("destination", "Diretório não encontrado")
	}
	if len(p.Sources) == 0 {
		ve.AddMessage("sources", "Campo obrigatório")
	}
	for _, src := range p.Sources {
		if !utils.DirectoryExists(src) && !utils.FileExists(src) {
			ve.AddMessage("sources", "Fonte não encontrada")
			break
		}
	}
	cf := config.GetConfig()
	path := filepath.Join(cf.IpedProfileFolder, p.Profile)
	if !utils.DirectoryExists(path) {
		ve.AddMessage("profile", "Perfil não encontrado")
	}
}
