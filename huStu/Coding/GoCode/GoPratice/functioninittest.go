package main

import "fmt"
//初始化顺序:变量初始化->init()->main()

var i int = initVar()

func init() {
	fmt.Println("init1.....")
}


func initVar() int {
	fmt.Println("initVar......")
	return 100
}


func main() {
	fmt.Println("main.....")
}