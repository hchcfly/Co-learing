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
	// 对磁盘对象进行定位
	locate.CollectObjects()
	//  心跳包
	go heartbeat.StartHeartbeat()
	//  定位文件是否存在
	go locate.StartLocate()
	http.HandleFunc("/objects/",objects.Handler)
	http.HandleFunc("/temp/",temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"),nil))
}