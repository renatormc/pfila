package processes

import (
	"fmt"
	"os/exec"

	"github.com/reantormc/pfila/api/database/models"
)

type Params interface {
	ToCmd() *exec.Cmd
}

func GetCmd(proc *models.Process) (*exec.Cmd, error) {
	var pars Params
	switch proc.Type {
	case "iped":
		pars = &IpedParams{}
	case "ftkimager":
		pars = &FtkimagerParams{}
	default:
		return nil, fmt.Errorf("type %q unknown", proc.Type)
	}
	if err := proc.GetParams(pars); err != nil {
		return nil, err
	}
	return pars.ToCmd(), nil
}
