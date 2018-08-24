package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello!"
	fmt.Printf("str1: %v:%T\n", str1, str1)

	var str2 string = "Hello!"
	fmt.Printf("str1: %v:%T\n", str2, str2)

	// Uppercase
	fmt.Println(strings.ToUpper(str1))
	// Title
	fmt.Println(strings.Title(str1))

	// Comparision
	lValue := "hello"
	uValue := "Hello"
	fmt.Println("Equal?", (lValue == uValue))
	fmt.Println("Equal non-case sensitive?", strings.EqualFold(lValue, uValue))

	fmt.Println("Contains exp?", strings.Contains(str1, "o"))
}
