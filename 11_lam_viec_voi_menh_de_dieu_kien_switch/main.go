package main

import (
	"fmt"
	// package math/rang để tạo các giá trị ngẫu nhiên
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Lam viec voi menh de switch")

	// tạo giá trị ngẫu nhiên từ dãy số thời gian
	rand.Seed(time.Now().Unix())

	// biến down nhận giá trị số dương
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string
	switch dow {
	case 1:
		result = "It's Sunday!"
	case 2:
		result = "It's Monday!"
		// nếu bạn dùng fallthrough sau khi thực hiện đoạn code trong case 2
		// sẽ thực hiện đoạn code bên dưới là default
		//fallthrough
	default:
		result = "It's some other day!"
	}
	fmt.Println(result)
}