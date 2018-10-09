package main

import "fmt"

func main() {

	sum := 1
	fmt.Println("Sum:", sum)

	colours := []string{"Red", "Green", "Blue"}
	fmt.Println(colours)

	sum = 0
	for i := 0; i < 10; i++ {
		sum = sum + i // or sum += i
	}
	fmt.Println("Sum:", sum)

	for i := 0; i < len(colours); i++ {
		fmt.Println(colours[i])
	}

	// Using range
	for i := range colours {
		fmt.Println(colours[i])
	}

	// Alternate to while loop
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto endofprogram
		}
		if sum > 500 {
			break // breaks the loop
		}
	}

	endofprogram : fmt.Println("end of program")
}
