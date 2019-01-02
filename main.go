package main

import (
	"log"
	"net/http"
	"./views"
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/disk", views.DiskHandler)
	mux.HandleFunc("/memory", views.MemoryHandler)
	mux.HandleFunc("/cpu", views.CpuHandler)
	log.Fatal(http.ListenAndServe(":8000", mux))
}