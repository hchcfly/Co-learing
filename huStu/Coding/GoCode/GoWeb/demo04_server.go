package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func login(w http.ResponseWriter,r *http.Request) {
	fmt.Printf("Method: %v\n", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token  := fmt.Sprintf("%x",h.Sum(nil))
		fmt.Println("token---> ",token)
		t, _ := template.ParseFiles("test.gtpl")
		t.Execute(w,token)
	} else {
		//请求的是登录数据,进行登录判断
		r.ParseForm()  //解析表单数据
		token := r.Form.Get("token")
		if token != "" {
			fmt.Printf("token: %v\n", token)
		} else {
			fmt.Println("token error")
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
        fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
        fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
        template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}

}


func main() {
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServer: ",err)
	}
}