package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond*100)
	}
}

func main() {
	go showMsg("c++") //启动了一个协程  goroutine1
	go showMsg("golang")      //goroutine2
	time.Sleep(time.Millisecond*2000)
	fmt.Println("main end...")  
}