package main

import "fmt"

func main() {
	str1 := "Hello!"
	str2 := "my name is"
	str3 := "Akshay."
	isTrue := true
	num := 50

	// Printing all
	fmt.Println(str1, str2, str3) // adds space between each word automatically with new line
	fmt.Print(str1)               // Print without new line.

	// String concatenation
	fmt.Println(str1 + str2 + str3) // Prints with no spaces with new line

	// Printing boolean
	fmt.Println(isTrue)

	// Formatted printing
	fmt.Printf("Hello %s", str3)

	// With error
	stringLength, err := fmt.Println(str1, str2)
	if err == nil {
		fmt.Println("String length is", stringLength)
		fmt.Println(err) // you don't have to print this. Just from information.
	}

	fmt.Printf("The number is %v\n", num)
	fmt.Printf("isTrue is %v\n", isTrue)
	fmt.Printf("float of the num is %.2f\n", float64(num))

	// Prints out the format.
	fmt.Printf("Type: %T, %T, %T, %T and %T\n", str1, str2, str3, isTrue, num)
	dataTypes := fmt.Sprintf("Type var: %T, %T, %T, %T and %T\n", str1, str2, str3, isTrue, num)
	fmt.Println(dataTypes)
}
