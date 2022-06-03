package main

import "fmt"

func main() {

	// 声明变量
	// var identifier type

	/*  var name string
	var age int
	var gender bool
	*/
	//批量声明
	var (
		name2   string
		age2    int
		gender2 bool
	)
	fmt.Printf("name2: %v\n", name2)
	fmt.Printf("age2: %v\n", age2)
	fmt.Printf("gender2: %v\n", gender2)

	//变量初始化
	//var 变量名 类型 = 表达式
	var name3 string = "李华"
	var age3 int = 18
	var gender3 = true
	fmt.Printf("name3: %v\n", name3)
	fmt.Printf("age3: %v\n", age3)
	fmt.Printf("gender3: %v\n", gender3)

	//类型推导
	var name4 = "lee"
	var age4 = 10
	var gender4 = false
	fmt.Printf("name4: %v\n", name4)
	fmt.Printf("age4: %v\n", age4)
	fmt.Printf("gender4: %v\n", gender4)

	//初始化多个变量
	var name5, age5, gender5 = "lili", 20, true

	fmt.Printf("name5: %v\n", name5)
	fmt.Printf("age5: %v\n", age5)
	fmt.Printf("gender5: %v\n", gender5)

	//短变量声明
	name := "张三"
	age := 15
	gender := false
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("gender: %v\n", gender)

	//匿名变量

}
