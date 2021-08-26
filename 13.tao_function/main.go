package main

import "fmt"

func main() {
	fmt.Println("functions")
	// gọi tới function tên là doSomething
	doSomething()

	// khai báo biến sum kiệu int nhận giá trị của functionaddValues với 2 params
	sum := addValues(3,4)
	fmt.Println("The sum is", sum)

	// khai báo 2 biến multiSum,, multiCount giá trị của function addAllValues với 3 params
	multiSum, multiCount := addAllValues(3,4,5)
	fmt.Println("The sum is", multiSum)
	fmt.Println("Count of items is", multiCount)
}

// function tên là doSomething
func doSomething()  {
	fmt.Println("Doing something")
}

// function tên là addValues với 2 params value1, value2
func addValues(value1 int, valu2 int) int {
	return value1 + valu2
}

// function tên là addAllValues với 1 params là values
func addAllValues(values ...int) (int,int) {
	total := 0
	// không dùng khóa nên khai báo tên biến là _
	for _,v := range values {
		total += v
	}
	// trả về 2 biến total,length of values
	return total,len(values)
}