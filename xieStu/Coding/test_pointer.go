package main

import "fmt"

func main() {

	var a int = 1
	var b = "hello"
	c := 3.14 //默认float64

	var p0 *int
	var p1 *string
	//var p2 *float32
	var p2 *float64
	p0 = &a
	p1 = &b
	p2 = &c

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)

	fmt.Printf("*p0: %v\n", *p0)
	fmt.Printf("*p1: %v\n", *p1)
	fmt.Printf("*p2: %v\n", *p2)

	fmt.Printf("p0: %v\n", p0)
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)

}
