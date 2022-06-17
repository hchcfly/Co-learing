package main

import (
	"./heartbeat"
	"./objects"
	"./locate"
	"net/http"
	"os"
	"log"
)


func main() {
	//  心跳包
	go heartbeat.StartHeartbeat()
	//  定位文件是否存在
	go locate.StartLocate()
	http.HandleFunc("/objects/",objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"),nil))
}