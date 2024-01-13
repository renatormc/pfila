package processes

import (
	"strings"

	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/helpers"
)

type FreecmdParams struct {
	Cmd string `json:"cmd"`
}

func (FreecmdParams) IsDocker() bool {
	return false
}

func (p *FreecmdParams) ToCmdArgs(proc *models.Process) ([]string, error) {
	args := strings.Fields(p.Cmd)
	return args, nil
}

func (p *FreecmdParams) Validate(ve *helpers.ValidationError) {

}
