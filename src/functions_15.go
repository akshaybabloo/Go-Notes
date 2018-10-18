package main

import "fmt"

func main() {
	doSomething()
}

func doSomething() {
	fmt.Println("Doing something!")
	fmt.Println(addValue(10, 20))

	sum := addAllValues(10, 20, 30)
	fmt.Println("Sum: ", sum)
}

// Takes in two integers, adds them and returns an integer.
func addValue(value1 int, value2 int) int {
	return value1 + value2
}

// Arbitrary function. Always takes in slice.
func addAllValues(values ...int) int {
	sum := 0
	for i := range values{
		sum += values[i]
	}
	fmt.Printf("%T\n", values) // Slice of integer value.
	return sum
}
