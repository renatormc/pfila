package processes

import (
	"os/exec"
)

type RunningCmds struct {
	Cmds map[uint]*exec.Cmd
}

func (rc *RunningCmds) SaveRunningCmd(id uint, cmd *exec.Cmd) {
	rc.Cmds[id] = cmd
}

func (rc *RunningCmds) GetRunningCmd(id uint) *exec.Cmd {
	cmd, ok := rc.Cmds[id]
	if !ok {
		return nil
	}
	return cmd
}

func (rc *RunningCmds) DeleteRunningCmd(id uint) {
	cmd := rc.GetRunningCmd(id)
	if cmd != nil {
		delete(rc.Cmds, id)
	}
}

func (rc *RunningCmds) GetIDs() []uint {
	var ids []uint
	for key := range rc.Cmds {
		ids = append(ids, key)
	}
	return ids
}

var runningCmds *RunningCmds

func GetRunningCmds() *RunningCmds {
	if runningCmds == nil {
		runningCmds = &RunningCmds{
			Cmds: make(map[uint]*exec.Cmd),
		}
	}
	return runningCmds
}
