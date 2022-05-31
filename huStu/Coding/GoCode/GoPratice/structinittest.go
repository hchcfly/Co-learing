package main

import "fmt"

func main() {
	type Person struct {
		id int
		name string
		age int
		email string
	}
	var tomp *Person
	var tom Person
	tomp = &tom
	fmt.Printf("tom: %v\n", tom)
	tom = Person{
		id:101,
		name:"tom",
		age:20,
		email:"hc@163.com",
	}
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("tomp: %v\n", tomp)
}