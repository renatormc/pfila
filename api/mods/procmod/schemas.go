package procmod

import (
	"log"
	"strings"

	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/helpers"
	"github.com/renatormc/pfila/api/processes"

	"github.com/renatormc/pfila/api/utils"
)

type ProcSchemaDump struct {
	ID           uint   `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	User         string `json:"user"`
	Pid          int    `json:"pid"`
	CreatedAt    string `json:"created_at"`
	Start        string `json:"start"`
	StartWaiting string `json:"start_waiting"`
	Finish       string `json:"finish"`
	Status       string `json:"status"`
	RandomID     string `json:"random_id"`
	Dependencies string `json:"dependencies"`
	Params       string `json:"params"`
}

func SerializeProc(p *models.Process) ProcSchemaDump {
	schema := ProcSchemaDump{
		ID:           p.ID,
		Name:         p.Name,
		Type:         p.Type,
		User:         p.User,
		Status:       p.Status,
		Params:       p.Params,
		CreatedAt:    helpers.SerializeTime(p.CreatedAt),
		Start:        helpers.SerializeTime(p.Start),
		StartWaiting: helpers.SerializeTime(p.StartWaiting),
		Finish:       helpers.SerializeTime(p.Finish),
		RandomID:     p.RandomID,
		Pid:          p.Pid,
		Dependencies: utils.SplitToString(p.GetDependencies(), ","),
	}
	return schema
}

func SerializeManyProc(mds []models.Process) []ProcSchemaDump {
	resp := make([]ProcSchemaDump, 0, len(mds))
	for _, m := range mds {
		schema := SerializeProc(&m)
		resp = append(resp, schema)
	}
	return resp
}

type ProcSchemaLoad struct {
	Type         string `json:"type"`
	Name         string `json:"name"`
	User         string `json:"user"`
	Dependencies string `json:"dependencies"`
	Params       string `json:"params"`
}

func (pl *ProcSchemaLoad) Fill(m *models.Process) *helpers.ValidationError {
	ve := helpers.NewValidationError()
	m.Name = strings.TrimSpace(pl.Name)
	if m.Name == "" {
		ve.AddMessage("name", "Campo obrigatório")
	}
	m.Type = pl.Type
	if m.Type != "iped" && m.Type != "ftkimager" {
		ve.AddMessage("type", "Tipo não conhecido")
	}
	m.User = strings.TrimSpace(pl.User)
	if m.User == "" {
		ve.AddMessage("user", "Campo obrigatório")
	}

	text := strings.TrimSpace(pl.Dependencies)
	if text != "" {
		parts := strings.Split(pl.Dependencies, ",")
		vals, err := utils.StringSlice2UintSlice(parts)
		if err != nil {
			ve.AddMessage("dependencies", "Valor incorreto")
		} else {
			m.SetDependencies(vals)
		}
	} else {
		m.SetDependencies([]uint{})
	}

	m.Params = pl.Params
	pars, err := processes.GetParams(m)
	if err != nil {
		log.Println(err)
		ve.AddMessage("internal", "erro interno")
		return ve
	}
	m.IsDocker = pars.IsDocker()
	pars.Validate(ve)
	if len(ve.Messages) > 0 {
		return ve
	}
	return nil
}
