package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
)

func sayhello(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()  //解析表单数据
	fmt.Println(r.Form)
	fmt.Printf("r.URL.Path: %v\n", r.URL.Path)
	fmt.Printf("r.URL.Scheme: %v\n", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	v := url.Values{}
	v.Set("name","Ava")
	v.Add("friend","hchc1")
	v.Add("friend","hchc2")
	v.Add("friend","hchc3")

	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])

	for k, v := range r.Form {
		fmt.Printf("key: %v\n", k)
		fmt.Println("val:",strings.Join(v,""))

	}

	fmt.Fprintf(w,"Hello my route!")  //信息输出到客户端

}

func login(w http.ResponseWriter,r *http.Request) {
	fmt.Printf("r.Method: %v\n", r.Method)  //打印请求的方法
	r.ParseForm()  //解析表单数据
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")  //解析html文件,返回值结果为t
		t.Execute(w,nil)  //将已经解析的模板应用于指定的数据对象w
	} else {
		//请求的数据是登录数据
		fmt.Printf("username: %v\n", r.Form["username"])
		fmt.Printf("password: %v\n", r.Form["password"])
		fmt.Fprintf(w,"登陆成功~~~~~")  //信息输出到客户端
	}

}



func main() {
	http.HandleFunc("/hello",sayhello)  //设置访问的路由
	http.HandleFunc("/login",login)  //设置访问的路由
	err := http.ListenAndServe(":8080", nil)  //设置并启动监听端口
	if err != nil {
		log.Fatal("ListenAndServer: ",err)
	}
}