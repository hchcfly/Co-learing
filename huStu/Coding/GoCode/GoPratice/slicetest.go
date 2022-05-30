package main

import "fmt"

//import "fmt"

func f1() {

	var s1 []int
	var s2 []string
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

func f2() {
	var s2 = make([]int,2)
	fmt.Printf("s2: %v\n", s2)
}


func f3() {
	var s1 = []int{1,2,3}
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("cap(s1): %v\n", cap(s1))
	fmt.Printf("s1[0]: %v\n", s1[0])
}


func f4() {
	var s1 = []int{1,2,3}
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("cap(s1): %v\n", cap(s1))
	fmt.Printf("s1[0]: %v\n", s1[0])
}

func main() {
	//f1()
	f3()
}