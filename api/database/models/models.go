package models

import (
	"encoding/json"
	"time"
)

type Process struct {
	ID           uint `gorm:"primarykey"`
	Script       string
	User         string
	Pid          int
	Start        time.Time
	StartWaiting time.Time
	Finish       time.Time
	Status       string
	RandomID     string
}

func (Process) TableName() string {
	return "process"
}

func (c *Process) SetJson(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.Value = string(data)
	return nil
}

func (c *Process) ParseJson(v any) error {
	err := json.Unmarshal([]byte(c.Value), v)
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
