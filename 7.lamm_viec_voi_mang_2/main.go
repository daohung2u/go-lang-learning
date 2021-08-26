package main

import (
	"fmt"
	// package sort các phần tử trong mãng
	"sort"
)

func main() {
	fmt.Println("Lam viec voi map")

	// khai báo biến color kiểu mãng với số phần tử không xác định
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// hàm append dùng để thêm 1 phần tử vào 1 mãng
	colors = append(colors, "Purple")
	fmt.Println(colors)

	// colors[1:len(colors) thêm 1 phần tử vào mãng tại vị trí 1
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	// xóa 1 phần tử của mãng tại vị trí 3
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	// tạo mới 1 mãng có 5 phần tử
	numbers := make([]int,5)
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 90
	numbers[3] = 60
	// hiện tại không có phần tử 5, nó sẽ là 0
	fmt.Println(numbers)

	// sắp xếp mảng theo thứ tự tăng dần
	sort.Ints(numbers)
	fmt.Println(numbers)
}
