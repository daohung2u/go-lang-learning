package main

import (
	"fmt"
	// package io đùng để đọc ghi gile
	"io"
	// package io đùng để đọc ghi gile
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Viet va doc file")

	content := "Hello from Go"
	// khai báo biến files nhận giá trị từ function Create
	files, err := os.Create("./fromstring.txt")
	checkError(err)

	// khai báo biến length nhận giá trị từ function WriteString
	length, err := io.WriteString(files, content)
	checkError(err)

	fmt.Printf("Wrote a file with%v characters\n", length)
	// defer keyword dùng để thực thi method cho đến khi function return
	defer files.Close()
	defer readFile("./fromstring.txt")
}

func checkError(err error)   {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	// khai báo biến data nhận giá trị từ funciton ReadFile
	// biến err nhận giá trị err
	data, err := ioutil.ReadFile(filename)
	checkError(err)
	fmt.Println("Text from a file: ", data)
}