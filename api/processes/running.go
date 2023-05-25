package processes

import (
	"os/exec"
)

var runningCmd map[uint]*exec.Cmd

func initialize() {
	if runningCmd == nil {
		runningCmd = make(map[uint]*exec.Cmd)
	}
}

func SaveRunningCmd(id uint, cmd *exec.Cmd) {
	initialize()
	runningCmd[id] = cmd
}

func GetRunningCmd(id uint) *exec.Cmd {
	initialize()
	cmd, ok := runningCmd[id]
	if !ok {
		return nil
	}
	return cmd
}

func DeleteRunningCmd(id uint) {
	initialize()
	cmd := GetRunningCmd(id)
	if cmd != nil {
		delete(runningCmd, id)
	}
}
