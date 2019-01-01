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
	log.Fatal(http.ListenAndServe(":8000", mux))
}