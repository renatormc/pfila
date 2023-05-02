package procmod

import (
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/helpers"
)

type ProcSchemaDump struct {
	ID           uint
	Type         string
	Name         string
	User         string
	Pid          int
	CreatedAt    string
	Start        string
	StartWaiting string
	Finish       string
	Status       string
	RandomID     string
	Params       string
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
		Start:        helpers.SerializeTime(p.CreatedAt),
		StartWaiting: helpers.SerializeTime(p.CreatedAt),
		Finish:       helpers.SerializeTime(p.CreatedAt),
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
