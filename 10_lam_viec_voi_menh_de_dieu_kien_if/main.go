package main

import "fmt"

func main() {
	fmt.Println("Dieu kien iff")
	theAnswer :=424
	var result string

	if theAnswer < 0 {
		result = "Less then zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater then zero"
	}

	fmt.Println(result)

	// khai báo biến và thực hiện điều kiện
	if x := -42; x < 0 {
		result = "Less then zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater then zero"
	}
	fmt.Println(result)
}
