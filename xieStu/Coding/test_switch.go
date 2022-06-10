package main

import (
	"fmt"
)

func f1(score int) {
	/* 	grade := 90
	   	switch grade {
	   	case 'A':
	   		fmt.Println("优秀")
	   	case 'B':
	   		fmt.Println("良好")
	   	case 'C':
	   		fmt.Println("及格")
	   	case 'D':
	   		fmt.Println("不及格")
	   	default:
	   		fmt.Println("其他")

	   	}
	*/
	switch {
	case score >= 90 && score <= 100:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("及格")
	case score < 60 && score > 0:
		fmt.Println("不及格")
	default:
		fmt.Println("成绩不合法")
	}

}

func main() {
	var score int
	fmt.Scan(&score)
	f1(score)
}
