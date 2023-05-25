package processes

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/external"
	"github.com/renatormc/pfila/api/helpers"

	"github.com/renatormc/pfila/api/utils"
)

type FtkimagerParams struct {
	Disk        string `json:"disk"`
	Destination string `json:"destination"`
	Verify      bool   `json:"verify"`
	Format      string `json:"format"`
}

func (FtkimagerParams) IsDocker() bool {
	return false
}

func (p *FtkimagerParams) ToCmdArgs(proc *models.Process) ([]string, error) {
	disks, err := external.GetDisks()
	if err != nil {
		return nil, err
	}
	if !utils.SliceContains(disks, p.Disk) {
		return nil, fmt.Errorf("disco %q não encontrado", p.Disk)
	}
	parts := strings.Split(p.Disk, " ")
	cf := config.GetConfig()
	args := []string{cf.Ftkimager, parts[0], p.Destination}
	if p.Format == "e01" {
		args = append(args, "--e01")
	}
	if p.Verify {
		args = append(args, "--verify")
	}
	return args, nil
}

func (p *FtkimagerParams) Validate(ve *helpers.ValidationError) {
	disks, err := external.GetDisks()
	if err != nil {
		log.Println(err)
		ve.AddMessage("internal", "internal error")
		return
	}
	if !utils.SliceContains(disks, p.Disk) {
		ve.AddMessage("disk", "Disco não encontrado")
	}
	dest := strings.TrimSpace(p.Destination)
	if dest == "" {
		ve.AddMessage("destination", "Campo obrigatório")
	} else {
		folder := filepath.Dir(dest)
		if !filepath.IsAbs(folder) || !utils.DirectoryExists(folder) {
			ve.AddMessage("destination", "Caminho inválido")
		}
	}

}
