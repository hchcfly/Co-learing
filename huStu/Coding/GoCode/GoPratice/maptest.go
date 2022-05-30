package main

import "fmt"


func test1() {
	//declare
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m1: %T\n", m1)
}

func test2() {
	var m1 = map[string]string{"name":"tom","age":"20","email":"@163.com"}
	fmt.Printf("m1: %v\n", m1)

	m2 := make(map[string]string)
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["email"] = "163.com"
	fmt.Printf("m2: %v\n", m2)
}


func test3() {
	var m1 = map[string]string{"name":"tom","age":"20","email":"@163.com"}
	var key = "name"
	var value = m1[key]
	fmt.Printf("value: %v\n", value)
}


func test4() {
	var m1 = map[string]string{"name":"tom","age":"20","email":"@163.com"}
	var k1 = "name"
	var k2 = "age1"
	v,ok := m1[k1]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	fmt.Println("----------------")
	v,ok = m1[k2]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	
}

func main() {
	test4()
}