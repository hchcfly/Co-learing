package main

import "fmt"

type USB interface {
	read()
	write()
}

type Computer struct {
	name string
}

func (c Computer) read() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("read.....")
}

func (c Computer) write() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("write....")
}

func main() {
	c := Computer{
		name:"联想",
	}
	c.read()
	c.write()
}