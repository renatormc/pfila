package models

import (
	"encoding/json"
	"time"
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

type Dependency struct {
	ID        uint `gorm:"primarykey"`
	BlockedID uint
	BlockerID uint
}

func (Dependency) TableName() string {
	return "dependency"
}
