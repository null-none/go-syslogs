package views

import (
	"net/http"
	"fmt"
	"../structure"
)

func MemoryHandler(w http.ResponseWriter, r *http.Request) {
	memory := structure.MemoryUsage()
	fmt.Fprintf(w, "Alloc: %v MiB\n", memory.Alloc)
	fmt.Fprintf(w, "TotalAlloc: %v MiB\n", memory.TotalAlloc)
	fmt.Fprintf(w, "Sys: %v MiB\n", memory.Sys)
}
