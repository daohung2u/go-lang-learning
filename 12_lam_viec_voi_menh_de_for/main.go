package main

import "fmt"

func main() {
	fmt.Println("Lam viec voi menh de for")
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// khai báo biến i kiêu số
	// i < len(colors)
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// khai báo biến i là kiểu số
	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum >= 200 {
			// goto dùng để di chuyển tới label trong chương trình
			goto theEnd
		}
	}

	// theEnd là label trong go lang
	theEnd:
	fmt.Println("End of Program")
}
