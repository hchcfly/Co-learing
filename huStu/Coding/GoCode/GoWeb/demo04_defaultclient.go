package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	//  建立request请求
	request, err := http.NewRequest(http.MethodGet, "http://www.chaindesk.cn/", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//  提交请求
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	b, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b: %s\n", b)
}