package main

import "fmt"

func main() {
	var pointer *int // nil pointer

	if pointer != nil {
		fmt.Println("Value of pointer: ", *pointer)
	} else {
		fmt.Println("The value of pointer is 'nil'")
	}

	var someNumber = 60
	pointer = &someNumber // Address of 'someNumber'
	fmt.Println("The value of pointer: ", *pointer)

	var pi float64 = 3.141
	pointer1 := &pi
	fmt.Println("Address of Pi: ", &pointer1)
	fmt.Println("Pi: ", *pointer1)

	*pointer1 = *pointer1 * 10
	fmt.Println("New Pi: ", *pointer1)
	fmt.Println("Actual Pi: ", pi)


}
