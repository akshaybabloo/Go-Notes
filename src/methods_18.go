package main

import "fmt"

type Dog struct {
	Breed string
	Weight int
	Sound string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d *Dog) SPeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{"Poodle", 30, "woof!"}
	fmt.Println(poodle)
	poodle.Speak()

	poodle.Sound = "Arf!"
	poodle.Speak()

	poodle.SPeakThreeTimes()
	poodle.SPeakThreeTimes() // this will add some more content to it
}
