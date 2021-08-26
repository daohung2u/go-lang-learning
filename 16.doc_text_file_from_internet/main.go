package main

import (
	"fmt"
	"io/ioutil"
	// package net/http cung cấp HTTP clients và triển khai server
	"net/http"
)

// khai báo url cần lấy dữ liệu
const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Doc text file from internet")

	// khai báo biến resp nhận dữ liệu từ function Get
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n",resp)
	// chờ cho đến khi function này thực thi xong
	defer resp.Body.Close()

	// khai báo biến bytes nhận giá trị từ ReadAll
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// khai báo biến content nhận giá trị từ function string(ép kiểu utf-8)
	content := string(bytes)
	fmt.Println(content)
}