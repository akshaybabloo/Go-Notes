package main

import (
	"fmt"
	"sort"
)

func main() {

	var colour = []string{"red", "green", "blue"} // A colour of strings
	fmt.Println(colour)

	colour = append(colour, "yellow") // Adding to an existing slice
	fmt.Println(colour)

	// To remove the first item
	colour = append(colour[1:len(colour)]) // You can also add colour[1:]
	fmt.Println(colour)

	// To remove the last item
	colour = append(colour[:len(colour)-1])
	fmt.Println(colour)

	// To have an initial size of the slice you can use make
	numbers := make([]int, 5, 5) // make(type of item, initial length, capping length)
	numbers[0] = 1
	numbers[1] = 10
	numbers[2] = 100
	numbers[3] = 1000
	numbers[4] = 10000
	fmt.Println(numbers)

	// You can also append an item to make and check its capacity
	numbers = append(numbers, 400)
	fmt.Println(numbers)
	fmt.Println(cap(numbers)) // This will not be 6, but a different number with more memory so that you can add things to it later.

	// Sorting slices.
	// Ascending order
	sort.Ints(numbers)
	fmt.Println(numbers)
}
