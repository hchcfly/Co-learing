package main

import "fmt"

func main() {
	a := [3]int{1,2,3}
	var pa [3]*int
	//指针 8byte
	fmt.Printf("pa: %v\n", pa)
	for i := 0; i < len(a); i++ {
		pa[i] = &a[i]
	}	

	fmt.Printf("pa: %v\n", pa)
}