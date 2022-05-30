package main

import "fmt"
import "strings"



func main() {
	// var s string = "hello world"
	// var s1 = "hello world"
	// s3 := "hello world"

	// // ` 
	// s4 := `
	// line 1
	// line 2
	// line 3
	// `

	// fmt.Printf("s: %v\n", s)
	// fmt.Printf("s1: %v\n", s1)
	// fmt.Printf("s3: %v\n", s3)
	// fmt.Printf("s4: %v\n", s4)

	// s1 := "tom"
	// s2 := "20"
	// msg := s1+s2
	// fmt.Printf("msg: %v\n", msg)

	// s1 := "tom"
	// s2 := "20"
	// msg := fmt.Sprintf("%s,%s",s1,s2)
	// fmt.Printf("msg: %v\n", msg)

	name := "Tom"
	age := "20"
	msg := strings.Join([]string{name,age},",")
	fmt.Printf("msg: %v\n",msg)


}