package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("file.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}

	// Custom errors
	myError := errors.New("new error")
	fmt.Println(myError)

	attendance := map[string]bool{
		"AAA": true,
		"BBB": true}
	attended, ok := attendance["AAA"] // Also know as - comma ok
	if ok {
		fmt.Println("AAA attended?", attended)
	} else {
		fmt.Println("No info for AAA.")
	}

	attended1, ok1 := attendance["CCC"]
	if ok1 {
		fmt.Println("AAA attended?", attended1)
	} else {
		fmt.Println("No info for CCC.")
	}

}
