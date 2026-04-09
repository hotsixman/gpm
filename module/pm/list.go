package pm

import (
	"geep/module/types"
	"strconv"
)

func (pm *PM) List() []types.ListElement {
	list := make([]types.ListElement, 0)
	for _, process := range pm.processArr {
		Pid := ""
		if process.cmd != nil {
			Pid = strconv.Itoa(process.cmd.Process.Pid)
		}

		elem := types.ListElement{
			Name:       process.name,
			Status:     string(process.status),
			CPUPercent: 0,
			Mem:        0,
			Recovered:  process.recoveredCount,
			Pid:        Pid,
		}

		if process.util != nil {
			cpuPercent, err := process.util.CPUPercent()
			if err == nil {
				elem.CPUPercent = cpuPercent
			}
			memInfo, err := process.util.MemoryInfo()
			if err == nil {
				elem.Mem = float64(memInfo.RSS) / 1024 / 1024
			}
		}

		list = append(list, elem)
	}
	return list
}
