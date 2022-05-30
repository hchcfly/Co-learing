package main

import "fmt"

func test1() {
	var a1 [2]int
	var a2 [3]string
	var a3 [3]bool
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	if a2[0] == "" {
		fmt.Println("no kg")
	} else {
		fmt.Println("yes")
	}
}


func main() {
	test1()
}