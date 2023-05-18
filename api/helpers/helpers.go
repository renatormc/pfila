package helpers

import (
	"errors"
	"math"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/process"
)

var kernel32 = syscall.MustLoadDLL("kernel32.dll")
var getProcessTimes = kernel32.MustFindProc("GetProcessTimes")

type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

func SerializeTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02/01/2006 15:04:05")
}

func GetProcess(pid int32, startTime time.Time) (*process.Process, error) {
	p, err := process.NewProcess(pid)
	if err != nil {
		return nil, err
	}

	createTime, err := p.CreateTime()
	if err != nil {
		return nil, err
	}
	ct := time.Unix(int64(createTime/1000), 0)
	if math.Abs(startTime.Sub(ct).Seconds()) < 30 {
		return p, nil
	}
	return nil, errors.New("not found")
}
