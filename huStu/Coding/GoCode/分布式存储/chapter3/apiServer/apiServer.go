package main

import (
	"./heartbeat"
	"./objects"
	"./locate"
	"./version"
	"net/http"
	"log"
	"os"
)


func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/",objects.Handler)
	http.HandleFunc("/locate/",locate.Handler)
	http.HandleFunc("/versions/",versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"),nil))
}