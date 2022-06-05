package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func login(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	fmt.Printf("username: %v\n", template.HTMLEscapeString(username))
	fmt.Printf("password: %v\n", template.HTMLEscapeString(r.Form.Get("password")))
	template.HTMLEscape(w,[]byte(username))
}

func main() {
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer",err)
	}
}