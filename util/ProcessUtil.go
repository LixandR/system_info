package util

import (
	"github.com/shirou/gopsutil/process"
	"strings"
)

func GetPidByName(pidName string) int32 {
	var pid int32 = -1
	pids, _ := process.Pids()
	for i := range pids {
		i2 := pids[i]
		newProcess, err := process.NewProcess(i2)
		if err==nil {
			name, err := newProcess.Name()
			if err == nil {
				if strings.EqualFold(name, pidName) {
					pid = i2
					break
				}
			}
		}
	}
	return pid
}