package models

import "time"

type Process struct {
	ID           uint `gorm:"primarykey"`
	Script       string
	User         string
	Pid          int64
	Start        time.Time
	StartWaiting time.Time
	Finish       time.Time
	Status       string
	Console      string
}

func (Process) TableName() string {
	return "process"
}

type Dependency struct {
	ID        uint `gorm:"primarykey"`
	BlockedID uint
	BlockerID uint
}

func (Dependency) TableName() string {
	return "dependency"
}
