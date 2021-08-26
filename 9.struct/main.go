package main

import "fmt"

func main() {
	fmt.Println("Struct")

	// khai báo biến poodle kiểu dữ liệu Dog
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeights: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeights: %v\n",poodle.Breed, poodle.Weight)
}

// Dog is a struct
type Dog struct {
	Breed string
	Weight int
}

