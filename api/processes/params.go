package processes

import (
	"fmt"

	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/helpers"
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
		pars = &IpedParams{}
	case "ftkimager":
		pars = &FtkimagerParams{}
	case "freecmd":
		pars = &FreecmdParams{}
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
