
package views

import (
	"net/http"
	"fmt"
	"../structure"
)


func CpuHandler(w http.ResponseWriter, r *http.Request) {
	cpu := structure.CpuUsage()
	fmt.Fprintf(w, "Cores: %v\n", cpu.cores)
}