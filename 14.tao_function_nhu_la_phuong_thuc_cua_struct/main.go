package main

import "fmt"

func main() {
	fmt.Println("Tao function nhu la phuong thuc cua struct")

	poodle := Dog{"Poodle", 10, "Woof!"}
	fmt.Println(poodle)
	fmt.Printf("Breed: %v\nWeight %v\n", poodle.Breed, poodle.Weight)

	// gọi function Speak của Dog struct
	poodle.Speak()
	poodle.Sound = "Arf!"
	poodle.Speak()

	// gọi function SpeakThreeTimes của Dog struct
	poodle.SpeakThreeTimes()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak function được thêm vào struct
func (d Dog) Speak()  {
	fmt.Println(d.Sound)
}

// SpeakThreeTimes function được thêm vào struct
func (d Dog) SpeakThreeTimes()  {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}