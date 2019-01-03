package views

import (
	"net/http"
	"fmt"
	"../structure"
)

func MemoryHandler(w http.ResponseWriter, r *http.Request) {
	memory := structure.MemoryUsage()
	fmt.Fprintf(w, "%+v", memory)
}
