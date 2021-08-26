package main

import (
	// package encoding/json dùng để xử lý json
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

	// khai báo biến tours nhận giá trị từ function tourFromJSon
	tours := tourFromJSon(content)
	for _,v := range tours {
		fmt.Println(v.Name)
	}
}

// dùng để chuyển đổi json thành các array tour
func tourFromJSon(content string) []Tour{
	// khai báo biến tours khởi tạo vùng nhớ cho 20 tour
	tours := make([]Tour,0,20)

	// khai báo biến decoder nhận giá trị của function NewDecoder
	decoder := json.NewDecoder(strings.NewReader(content))

	// khai báo biến _ nhận giá trị nil khi decoder xibg
	_,err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	// lặp cho đến khi còn nhận result
	for decoder.More() {
		// khai báo biên err nhận giá trị err từ function decode
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}

	return tours
}

// Tour struct
type Tour struct {
	Name, Price string
}