package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var str1 string = "Hello,world."
	var html string = `
	<html>
		<head><title>hello golang</title>
	<html>
	`
	fmt.Printf("str1: %v\n", str1)
	fmt.Printf("html: %v\n", html)

	//字符串连接
	name := "Tom"
	age := "20"
	msg := name + " " + age
	fmt.Printf("msg: %v\n", msg)

	msg = ""
	msg += name
	msg += " "
	msg += age
	fmt.Printf("msg: %v\n", msg)

	//fmt.Sprintf()函数
	msg = fmt.Sprintf("%s,%s", name, age)
	fmt.Printf("msg: %v\n", msg)

	//st(ring.Join()
	msg = strings.Join([]string{name, age}, ",")
	fmt.Printf("msg: %v\n", msg)

	// buffer.WriteString()
	var buffer bytes.Buffer
	buffer.WriteString("Tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String())

}
