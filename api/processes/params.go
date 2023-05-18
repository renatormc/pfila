package processes

import (
	"fmt"
	"os/exec"

	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/processes/ftkimager"
	"github.com/renatormc/pfila/api/processes/iped"
)

type Params interface {
	ToCmd() *exec.Cmd
}

func GetCmd(proc *models.Process) (*exec.Cmd, error) {
	var pars Params
	switch proc.Type {
	case "iped":
		pars = &iped.IpedParams{}
	case "ftkimager":
		pars = &ftkimager.FtkimagerParams{}
	default:
		return nil, fmt.Errorf("type %q unknown", proc.Type)
	}
	if err := proc.GetParams(pars); err != nil {
		return nil, err
	}
	return pars.ToCmd(), nil
}
