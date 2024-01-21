package processes

import (
	"runtime"
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
	var args []string
	if runtime.GOOS == "windows" {
		args = []string{"cmd", "/c"}
	} else {
		args = []string{"bash"}
	}

	args = append(args, strings.Fields(p.Cmd)...)
	return args, nil
}

func (p *FreecmdParams) Validate(ve *helpers.ValidationError) {

}
