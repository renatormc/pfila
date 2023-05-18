package helpers

import (
	"errors"
	"log"
	"math"
	"os"
	"time"

	"github.com/shirou/gopsutil/process"
)

func SerializeTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Local().Format("02/01/2006 15:04:05")
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

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Fatal(err)
	}

	return !info.IsDir()
}

func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
