package main

import (
	"fmt"
)

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i,%v\n", i)
		if i>= 5 {
			break;
		}
	}
}

//fallthrough 穿透
func test2() {
	i:=2
	switch i {
	case 1:
		fmt.Println("1")
		break;
	case 2:
		fmt.Println("2")
		//break
		fallthrough

	case 3:
		fmt.Println("3")
		break;
	}

}

//标签Label
func test3() {

MYLABEL:

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i>=5 {
			break MYLABEL
		}
	}

	fmt.Println("END...")
}

func main() {
	test3()
}