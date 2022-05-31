package main
import "fmt"

func f1() {
	fmt.Println("start.....")
	
	//延迟执行 栈的执行流程
	defer fmt.Println("step1.....")
	defer fmt.Println("step2.....")
	defer fmt.Println("step3.....")

	fmt.Println("end.......")
}

func main() {
	f1()
}