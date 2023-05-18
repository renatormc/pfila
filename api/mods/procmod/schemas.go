package procmod

import (
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/helpers"
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
	Type   string
	Name   string
	User   string
	Params string
}

func (pl *ProcSchemaLoad) Fill(m *models.Process) *helpers.ValidationError {
	m.Name = pl.Name
	m.Type = pl.Type
	m.User = pl.User
	m.Params = pl.Params
	return nil
}
