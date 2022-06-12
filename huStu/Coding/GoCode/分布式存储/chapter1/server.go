package main

import (
	"net/http"
	"log"
	"os"
	"./objects"
)


func main() {
	http.HandleFunc("/objects/",objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"),nil))
}