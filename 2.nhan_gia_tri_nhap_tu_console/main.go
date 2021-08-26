package main

import (
	// packgage hỗ trợ đọc viết ra vào console
	"bufio"
	"fmt"
	// package của hệ điều hành
	"os"
)

func main() {
	fmt.Println("Nhan gia tri nhap vao tu console")

	// khai báo biến reader kiểu dữ liệu Reader dùng để chờ dữ liệu được nhập từ console
	// os.Stdin nhận dữ liệu nhập vào hệ điều hành
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text:")

	// khai báo 2 biến nhận giá trị từ function ReadString
	// biết input nhận giá trị được nhập từ console
	// biến _ là biến error khi function ném ra error
	input, _ := reader.ReadString('\n')

	fmt.Println("You entered: ", input)
}
