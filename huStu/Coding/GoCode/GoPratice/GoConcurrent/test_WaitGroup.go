package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func showMsg(i int) {
	fmt.Printf("i: %v\n", i)
}



func main() {
	for i := 0; i < 10; i++ {
		go showMsg(i)
	}
	fmt.Println("end...")
}