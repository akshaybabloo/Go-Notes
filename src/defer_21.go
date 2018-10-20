package main

import "fmt"

func main() {
	defer fmt.Println("Deferred Statement 1")
	defer fmt.Println("Deferred Statement 2")

	NewFunction()

	defer fmt.Println("Deferred Statement 3")
	defer fmt.Println("Deferred Statement 4")
	fmt.Println("Un-deferred Statement")

	x := 1000
	defer fmt.Println("Value of x:", x)
	x++
	fmt.Println("value of x after increment:", x)
}

func NewFunction() {
	defer fmt.Println("Deferred in the function")
	fmt.Println("Not deferred in the function")
}
