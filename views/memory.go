package views

import (
	"runtime"
	"net/http"
    "fmt"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}


func MemoryHandler(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Fprintf(w, "Alloc: %v MiB\n", bToMb(m.Alloc))
	fmt.Fprintf(w, "TotalAlloc: %v MiB\n", bToMb(m.TotalAlloc))
	fmt.Fprintf(w, "Sys: %v MiB\n", bToMb(m.Sys))
}
