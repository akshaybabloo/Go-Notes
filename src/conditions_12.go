package main

import "fmt"

func main() {

	//var x float64 = 42 // Change this to see the result
	var result string

	if x := 42; x < 0 {
		result = "Less than zero"
	} else if x==0{
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println("Result:", result)
	
}
