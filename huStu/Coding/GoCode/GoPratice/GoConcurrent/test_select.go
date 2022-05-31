package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr =  make(chan string)




func main() {
	go func() {
		defer close(chanInt)
		defer close(chanStr)
		chanInt <- 100
		chanStr <- "hello"
	}()

	for {
		select {
		case r:=<-chanInt:
			fmt.Printf("rint: %v\n", r)
		case r:=<-chanStr:
			fmt.Printf("rstring: %v\n", r)
		default:
			fmt.Println("default....")
		}
		time.Sleep(time.Second)
	}
}