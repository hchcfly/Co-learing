package main

import "fmt"


func test2() {
	type Person struct {
		id int
		name string
		age int
	}
	tom := Person{
		id:101,
		name:"hchc",
		age:10,
	}
	var p_person *Person
	p_person = &tom
	fmt.Printf("p_person: %v\n", p_person)
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %v\n", *p_person)
}


func test3() {
	type Person struct {
		id int
		name string
		age int
	}
	//分配内存空间
	var tom = new(Person)
	fmt.Printf("tom: %v\n", tom)
}


func main() {
	test3()
}