package main

import (
	"bytes"
	"fmt"
	"strings"
	//"strings"
)



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

	//**********strings.Join()*******************//
	name := "Tom"
	age := "20"
	msg := strings.Join([]string{name,age},",")
	fmt.Printf("msg: %v\n",msg)

	//********buffer.WriteString()***************//
	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	//********切片******************************//

	s := "hello world"
	// n := 3
	// m := 5
	// fmt.Printf("len(str): %v\n",len(str))
	// fmt.Printf("str: %v\n", str[n])
	// fmt.Println(str[n])
	// fmt.Println(str[n:m])
	// fmt.Println(str[n:])
	// fmt.Println(str[:m])

	//*************Split****************************//
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))

	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))

	fmt.Printf("strings.ToLower(s): %v\n", strings.ToLower(s))

	fmt.Printf("strings.ToUpper(): %v\n", strings.ToUpper(s))
	//前缀
	fmt.Printf("strings.HasPrefix(\"hello\"): %v\n", strings.HasPrefix(s,"Hello"))
	//结尾
	fmt.Printf("strings.HasSuffix(\"world\"): %v\n", strings.HasSuffix(s,"world"))

	//查找index
	fmt.Printf("strings.Index(s, \"ll\"): %v\n", strings.Index(s, "ll"))
	fmt.Printf("strings.LastIndex(\"w\"): %v\n", strings.LastIndex(s,"w"))

}