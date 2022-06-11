package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)


func main() {
	http.HandleFunc("/objects/",objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"),nil))
}