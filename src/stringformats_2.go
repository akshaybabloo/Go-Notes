package main

import "fmt"

func main() {
	str1 := "Hello!"
	str2 := "my name is"
	str3 := "Akshay."
	str4 := true

	// Printing all
	fmt.Println(str1, str2, str3) // adds space between each word automatically with new line
	fmt.Print(str1) // Print without new line.

	// String concatenation
	fmt.Println(str1 + str2 + str3) // Prints with no spaces with new line

	// Printing boolean
	fmt.Println(str4)

	// Formatted printing
	fmt.Printf("Hello %s", str3)

	// With error
	stringLength, err := fmt.Println(str1, str2)
	if err == nil {
		fmt.Println("String length is", stringLength)
		fmt.Println(err) // you don't have to print this. Just from information.
	}

}
