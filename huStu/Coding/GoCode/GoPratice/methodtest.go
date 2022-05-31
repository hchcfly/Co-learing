package main

import "fmt"

type Person struct {
	name string
}

//eat属于Person结构体
func (per Person)eat() {
	fmt.Printf("%v,eating\n",per.name)
}

func (per Person)sleep() {
	fmt.Printf("%v,sleep\n",per.name)
}

type Customer struct {
	name string
}

func (customer Customer) login(name string,pwd string)bool {
	fmt.Printf("customer.name: %v\n", customer.name)
	if name == "tom" && pwd == "123" {
		return true
	} else {
		return false
	}
}

func main() {
	// per:=Person{
	// 	name:"tom",
	// }
	// per.eat()
	// per.sleep()

	customer := Customer{
		name:"tom",
	}

	b := customer.login("tom","123")
	fmt.Printf("b: %v\n", b)

}