package apiServer

import (
	"./heartbeat"
	"./objects"
	"./locate"
	"net/http"
	"log"
	"os"
)


func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/",objects.Handler)
	http.HandleFunc("/locate/",locate.Handler)
	log.Fatal(http.ListenAndServer(os.Getenv("LISTEN_ADDRESS"),nil))
}