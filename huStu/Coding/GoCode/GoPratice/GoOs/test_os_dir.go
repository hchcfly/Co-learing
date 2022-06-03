package main

import (
	"fmt"
	"os"
)

//官方帮助文档: https://pkg.go.dev/std

//  创建文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

//  创建目录
func makeDir() {
	err := os.Mkdir("a",os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	//创建多级目录
	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if err2 != nil {
		fmt.Printf("err: %v\n", err)
	}
}

//  删除目录或文件
func remove() {
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err2 := os.RemoveAll("a")
	if err2 != nil {
		fmt.Printf("err: %v\n", err)
	}
}

//  获取当前工作目录
func wd() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
	
	//修改当前目录:os.Chdir("")
}

//  重命名os.Rename()


//  读文件
func readFile() {
	b, _ := os.ReadFile("a.txt")
	fmt.Printf("b: %v\n", string(b[:]))
}

//  写文件
func writeFile() {
	os.WriteFile("a.txt", []byte("hello world"), os.ModePerm)

}


func main() {
	//createFile()
	//makeDir()
	//remove()
	readFile()
	writeFile()
}