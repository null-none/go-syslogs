package structure

import (
	"runtime"
)

type CpuStatus struct {
	Cores  uint64 `json:"cores"`
}
  
func CpuUsage() (cpu CpuStatus) {
	cpu.Cores = uint64(runtime.NumCPU())
	return cpu
 }
 