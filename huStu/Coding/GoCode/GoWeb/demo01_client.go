package main

import (
	"fmt"
	"net/http"
)



func main() {
	requestUrl := "http://www.baidu.com"
	r, err := http.Get(requestUrl)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer r.Body.Close()
	fmt.Println(r.Body)
}