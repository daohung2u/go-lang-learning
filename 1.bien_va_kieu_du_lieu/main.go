package main

import (
	// package dùng để tương tác với console, in ra console
	"fmt"
)

// khai báo hằng aConst không thể thay đổi được
const aConst int = 64

func main() {
	fmt.Println("Khai bao bien và kieu du lieu co ban")

	// - Kieu 1

	// --- khai báo biến aString kiểu dữ liêu string
	var aString string
	var aInt int

	// - kieu 2
	// --- khai báo biến aString2 kiểu dữ liệu string
	aString2 := ""

	// - Gán giá trị
	aString = "This is a string"
	aString2 = "This is a another string"
	aInt = 10

	// - In dữ liệu ra console
	fmt.Println(aString)
	fmt.Println(aString2)
	fmt.Println(aInt)

	// Các vd khác
	var aString3 string = "This is Go!"
	fmt.Println(aString3)

	// In ra console với format: %T là loại dữ liệu
	fmt.Printf("The variable's type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is a another string"
	fmt.Println(anotherString)
	fmt.Printf("The variables's type is %T\n\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variables's type is %T\n", anotherInt)

	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("The variables's type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("The variables's type is %T\n", aConst)
}
