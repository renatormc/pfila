package processes

import (
	"fmt"

	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/helpers"
	"github.com/renatormc/pfila/api/processes/ftkimager"
	"github.com/renatormc/pfila/api/processes/iped"
)

type Params interface {
	ToCmdArgs(*models.Process) ([]string, error)
	Validate(*helpers.ValidationError)
	IsDocker() bool
}

func GetParams(proc *models.Process) (Params, error) {
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
	return pars, nil
}

func GetCmdArgs(proc *models.Process) ([]string, error) {
	pars, err := GetParams(proc)
	if err != nil {
		return nil, err
	}
	return pars.ToCmdArgs(proc)
}
