
package views

import (
	"net/http"
	"fmt"
	"../structure"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func DiskHandler(w http.ResponseWriter, r *http.Request) {
	disk := structure.DiskUsage("/")
	fmt.Fprintf(w, "All: %.2f GB\n", float64(disk.All)/float64(GB))
	fmt.Fprintf(w, "Used: %.2f GB\n", float64(disk.Used)/float64(GB))
	fmt.Fprintf(w, "Free: %.2f GB\n", float64(disk.Free)/float64(GB))
}