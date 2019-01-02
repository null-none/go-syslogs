package structure

import (
	"runtime"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

type MemoryStatus struct {
	Alloc  uint64 `json:"alloc"`
	TotalAlloc uint64 `json:"total_alloc"`
	Sys uint64 `json:"sys"`
}
  
func MemoryUsage() (memory MemoryStatus) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memory.Alloc = bToMb(m.Alloc)
	memory.TotalAlloc =  bToMb(m.TotalAlloc)
	memory.Sys = bToMb(m.Sys)
	return memory
 }
 