package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "blue"
	colors[2] = "Yellow"
	fmt.Println(colors)    // [Red blue Yellow]
	fmt.Println(colors[1]) // blue
	fmt.Println("Number of colors: ", len(colors))

	var numbers = [3]int{1, 2, 3}
	fmt.Println(numbers) // [1 2 3]
	fmt.Println("Number of numbers: ", len(numbers))
}
