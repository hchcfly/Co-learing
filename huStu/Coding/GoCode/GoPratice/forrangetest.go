package main

import "fmt"


func f1() {
	//数组
	var a = [5]int{1,2,3,4,5}
	for _,v := range a {
		fmt.Printf("%v\n",v)
	}
}

func f2() {
	//切片
	var s = []int{1,2,3}
	for _,v := range s {
		fmt.Printf("v: %v\n", v)
	}
}


func f3() {
	//map
	m := make(map[string]string, 0)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "@163.com"
	for key,value := range m {
		fmt.Printf("%v : %v\n", key,value)
	}
}

func main() {
	//f1()
	//f2()
	f3()
}