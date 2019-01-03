
package views

import (
	"net/http"
	"fmt"
	"../structure"
)

func DiskHandler(w http.ResponseWriter, r *http.Request) {
	disk := structure.DiskUsage("/")
	fmt.Fprintf(w, "%+v", disk)
}