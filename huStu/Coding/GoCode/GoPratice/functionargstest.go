package main

import "fmt"



func sum(a int,b int)(int) {
	
	return a+b
}

func f1(a int) {
	//copy
	a = 100
}

//传slice
func f2(s []int) {
	s[0] = 1000
}

//传变长参数
func f3(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {

	f3(1,2,3,4,5)
	f3(3,4,5,6,7,8)

	// s := []int{1,2,3}
	// f2(s)
	// fmt.Printf("s: %v\n", s)
	// x := 200
	// f1(x)
	// fmt.Printf("x: %v\n", x)
}