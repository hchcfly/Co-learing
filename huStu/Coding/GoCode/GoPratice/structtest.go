package main

import "fmt"

//自定义一个类型
type Person struct {
	id int
	name string
	age int
	email string
}

type Customer struct {
	id,age int
	name,email string 
}

func main() {
	// var tom Person
	// tom.id = 101
	// tom.name = "hchc"
	// fmt.Printf("tom: %v\n", tom)


	// //匿名结构体
	// var hc struct {
	// 	id int
	// 	name string 
	// 	age int
	// }
	// fmt.Printf("hc: %v\n", hc)
	// var tom Person
	// fmt.Printf("tom: %v\n", tom)
	// var kite Customer
	// fmt.Printf("kite: %v\n", kite)
}