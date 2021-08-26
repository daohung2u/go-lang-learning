package main

import "fmt"

func main() {
	fmt.Println("Lam viec voi mang")

	// khai báo biến colors kiểu mảng có 3 phần tử kiểu string
	// [3] số phẩn tử trong mãng được khai báo
	// string dữ liệu của mãng
	var colors[3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println( colors)
	fmt.Println( colors[0])

	// khai báo biến numbers kiểu mãng có 5 phần tử
	// {1,2,3,4,5} các phần tử trong mãng
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	fmt.Println("Number of colors: ", len(colors))
	fmt.Println("Number of numbers: ", len(numbers))
}