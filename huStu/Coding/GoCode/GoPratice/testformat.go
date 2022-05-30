package main

import "fmt"

type WebSite struct {
	Name string
}



func main() {
	site := WebSite{Name:"douke360"}
	fmt.Printf("site: %v\n", site)
	fmt.Printf("site: %#v\n", site)
	fmt.Printf("site: %T\n", site)  //打印出类型
	

}