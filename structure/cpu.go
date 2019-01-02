package structure

import (
	"runtime"
)

type CpuStatus struct {
	cores  int `json:"cpu"`
}
  
func CpuUsage() (cpu CpuStatus) {
	cpu.cores = runtime.NumCPU()
	return cpu
 }
 