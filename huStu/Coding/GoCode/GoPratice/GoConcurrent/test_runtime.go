package main

import (
	"fmt"
	"runtime"
)

func show(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}

func main() {
	go show("c++")

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("golang")
	}
}