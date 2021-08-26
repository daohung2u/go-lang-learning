package main

import (
	"fmt"
	// package time dùng để do lường và định dạng thời gian
	"time"
)

func main()  {
	// khai báo biến kiểu time
	n := time.Now()
	fmt.Println("I recored this video at", n)

	// khai báo biến t kiểu time
	t := time.Date(2009, time.November,10,23,22,0,0,time.UTC)

	fmt.Println(t.Format(time.ANSIC))

	// biến parsed kiểu time nhận từ function Parse 1 chuỗi kiểu time
	parsedTime, _ := time.Parse(time.ANSIC,"Tue Nov 10 23:00:00 2009")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)
}