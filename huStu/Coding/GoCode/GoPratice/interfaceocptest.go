package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	name string
}

type Cat struct {
	name string	
}

//Dog实现Pet接口
func (dog Dog)eat() {
	fmt.Println("Dog eat...")
}

func (dog Dog)sleep() {
	fmt.Println("Dog sleep...")
}

//Cat实现Pet接口
func (cat Cat)eat() {
	fmt.Println("Cat eat...")
}

func (cat Cat)sleep() {
	fmt.Println("Cat sleep...")
}

type Person struct {
	
}


func (person Person)care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	dog := Dog{
	}
	cat := Cat{
	}
	person := Person{
	}
	person.care(dog)
	person.care(cat)
}