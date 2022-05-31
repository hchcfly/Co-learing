package main

import (
	"fmt"
	"time"
)



func main() {
	// timer := time.NewTimer(time.Second*2)
	// fmt.Printf("time.Now(): %v\n", time.Now())
	// t1 := <-timer.C
	// fmt.Printf("t1: %v\n", t1)


	// fmt.Printf("time.Now(): %v\n", time.Now())
	// timer := time.NewTimer(time.Second*2)
	// <-timer.C
	// fmt.Printf("time.Now(): %v\n", time.Now())

	//创建一个一秒计时器
	timer := time.NewTimer(time.Second)
	go func(){
		//阻塞等待计时器到时间
		<-timer.C
		fmt.Println("func...")

	}()
	//结束定时器
	s := timer.Stop()
	if s {
		fmt.Println("stop...")
	}
	time.Sleep(time.Second*3)
	fmt.Println("main end...")
}