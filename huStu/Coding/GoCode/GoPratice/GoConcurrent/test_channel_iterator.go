package main

import "fmt"

//无缓冲
var c = make(chan int)


func main() {

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		v,ok := <-c
		if ok {
			fmt.Printf("v: %v\n", v)
		} else {
			break
		}
	}

	// for v := range c {
	// 	fmt.Printf("v: %v\n", v)
	// }

	// r := <-c
	// fmt.Printf("r: %v\n", r)
	// r = <-c
	// fmt.Printf("r: %v\n", r)
	// r = <-c
	// fmt.Printf("r: %v\n", r)
}