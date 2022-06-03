package main

import "fmt"


func main() {
	//const constantName [type] = value
	//编译阶段就确定下来的值，在程序运行时无法改变

	const PI float64 = 3.14
	const PI2 = 3.1415 //可以缺省类型
	fmt.Printf("PI: %v\n", PI)
	fmt.Printf("PI2: %v\n", PI2)
}
