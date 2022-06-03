package main

import "fmt"

func main() {
	var b bool = true
	var b2 bool = false

	var b3 = true
	var b4 = false

	b5 := true
	b6 := false

	fmt.Printf("b: %v\n", b)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("b3: %v\n", b3)
	fmt.Printf("b4: %v\n", b4)
	fmt.Printf("b5: %v\n", b5)
	fmt.Printf("b6: %v\n", b6)

	age := 18
	adult := age >= 18
	if adult {
		fmt.Printf("是成年人\n")
	} else {
		fmt.Printf("是未成年人\n")
	}

	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i: %v\n", i)
	}

	gender := "男"
	if age >= 18 && gender == "男" {
		fmt.Println("是成年男性")
	}
}
