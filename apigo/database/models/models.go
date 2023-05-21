package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/renatormc/pfila/api/utils"
)

type Process struct {
	ID           uint `gorm:"primarykey"`
	Type         string
	Name         string
	User         string
	Pid          int
	CreatedAt    time.Time
	Start        time.Time
	StartWaiting time.Time
	Finish       time.Time
	Status       string
	RandomID     string
	Params       string
	Dependencies string
	IsDocker     bool
}

func (Process) TableName() string {
	return "process"
}

func (proc *Process) SetParams(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	proc.Params = string(data)
	return nil
}

func (proc *Process) GetParams(v any) error {
	err := json.Unmarshal([]byte(proc.Params), v)
	if err != nil {
		return err
	}
	return nil
}

func (proc *Process) GetDependencies() []uint {
	ret := []uint{}
	if proc.Dependencies == "" {
		return ret
	}
	text := proc.Dependencies[1 : len(proc.Dependencies)-1]
	parts := strings.Split(text, ",")
	for _, p := range parts {
		v, err := strconv.ParseUint(p, 10, 32)
		if err == nil {
			ret = append(ret, uint(v))
		}
	}
	return ret
}

func (proc *Process) SetDependencies(deps []uint) {
	proc.Dependencies = fmt.Sprintf(",%s,", utils.SplitToString(deps, ","))
}
