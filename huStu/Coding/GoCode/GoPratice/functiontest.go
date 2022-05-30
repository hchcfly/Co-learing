package main

import "fmt"

func sum(a int,b int) (ret int) {
	ret = a+b
	return ret
}

func sum1(a int,b int) int {
	return a+b
}

func sum2()(name string,age int) {
	name = "tom"
	age = 20
	return
} 


func main() {
	//r := sum(1,2)
	//fmt.Printf("r: %v\n", r)
	fmt.Println(sum2())
}