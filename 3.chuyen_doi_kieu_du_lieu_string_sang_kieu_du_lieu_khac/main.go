package main

import (
	"bufio"
	"fmt"
	"os"
	// package chứa function để convert từ kiểu string sang kiểu dữ liệu khác
	"strconv"
	// làm việc với kiểu dữ liệu strings
	"strings"
)

func main() {
	fmt.Println("Chuyen du lieu tu string sang kieu du lieu khac")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text:")

	input, _ := reader.ReadString('\n')

	fmt.Println("You entered: ", input)

	fmt.Print("Enter a number:")

	numInput,_ := reader.ReadString('\n')

	// khai báo 2 biến nhận giá trị từ function ParseFloat
	// biết aFloat nhận giá trị khi chuyển đổi dữ liệu thành công
	// biến _ là biến error khi function ném ra error
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput),64)

	// xử lý khi error
	if err != nil {
		fmt.Println("Invalid number: ", err)
	} else {
		fmt.Println("Value of number: ", aFloat)
	}

}
