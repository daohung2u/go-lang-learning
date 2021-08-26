package main

import "fmt"

func main() {
	fmt.Println("Tham chieu value dung con tro")

	// khai báo biến P kiểu dữ liệu innt dùng contrỏ
	var p *int
	fmt.Println("Value of p:", p)

	anInt := 42
	// khai báo biến P2 kiểu dữ liệu innt dùng contrỏ
	var p2 = &anInt
	// *p2 dùng để lấy giá trị của con trỏ
	fmt.Println("Value of p2:", *p2)

	value1 := 42.13
	// khai báo biến contror pointer1 tham chiếu tới biến value1
	pointer1 := &value1
	fmt.Println("Value 1:", * pointer1)

	// gán giá trị cho biến contrỏ
	*pointer1 = *pointer1 / 31
	fmt.Println("Value 1:", value1)
	fmt.Println("Pointer 1:", *pointer1)
}
