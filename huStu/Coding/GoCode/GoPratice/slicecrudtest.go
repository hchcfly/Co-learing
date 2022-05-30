package main

import "fmt"

//add
func test1() {
	var s1 = []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	fmt.Printf("s1: %v\n", s1)
}

//del
func test2() {
	var s1 = []int{1,2,3,4}
	//i
	s1 = append(s1[:2],s1[3:]...)
	fmt.Printf("s1: %v\n", s1)
}

//浅拷贝
func test3() {
	var s1 = []int{1,2,3,4}
	var s2 = s1
	s2[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}


func main() {
	test3()
}