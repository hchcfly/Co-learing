package main

import "fmt"

type user struct {
	name string
}

func main() {
	u := user{"Tom"}
	fmt.Printf("% + v\n", u)
	fmt.Printf("%#v\n", u)
}
