package getCPU

import (
	"runtime"

	github "github.com/shirou/gopsutil/mem"
)

type CPUInfo struct {
	number      int
	total       uint64
	free        uint64
	usedPercent float64
}

func GetCPUinfo() CPUInfo {

	thisGithub, _ := github.VirtualMemory()

	CPUINFO := CPUInfo{}
	CPUINFO.number = runtime.NumCPU()
	CPUINFO.total = thisGithub.Total
	CPUINFO.free = thisGithub.Free
	CPUINFO.usedPercent = thisGithub.UsedPercent

	return CPUINFO
}
