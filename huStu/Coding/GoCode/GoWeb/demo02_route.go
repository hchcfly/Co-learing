package main

import (
	"fmt"
	"net/http"
)

//  自定义多路分发器
type MyMux struct {
	
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w,r)
		return
	}
	http.NotFound(w,r)
	return 
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello my route!")
}


func main() {
	mux := &MyMux{}
	http.ListenAndServe(":8080",mux)
}