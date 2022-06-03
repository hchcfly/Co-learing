package main

import "fmt"

func main() {
	var arr1 = [...]int{1, 2, 3, 4}
	for i, v := range arr1 {
		fmt.Printf("i:%v", i)
		fmt.Printf(" ")
		fmt.Printf("v:%v", v)
		fmt.Println("")
	}
}
