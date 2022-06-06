package main

import "fmt"

type Animal struct {
	name string
	age int
}

type Dog struct {
	Animal    //理解为inhiert
}


func (a *Animal) eat() {
	fmt.Println("eating...")
}

// func (a Animal) sleep() {
// 	fmt.Println("sleep...")
// }

func (d *Dog) get() {
	fmt.Println("i will get")
}

// type Dog struct {
// 	a     Animal    //理解为inhiert
// 	color string

// }


// type Cat struct {
// 	a Animal
// 	bbb string
// }


func main() {
	Dog := Dog{}
	Dog.get()
	
	// dog := Dog{
	// 	Animal{name:"hc",age:2},
	// 	"黑色",
	// }
	// dog.eat()
	// dog.sleep()
	// fmt.Printf("dog.color: %v\n", dog.color)
	// fmt.Printf("dog.a.name: %v\n", dog.name)
	// fmt.Printf("dog.a.age: %v\n", dog.age)
}