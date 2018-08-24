package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//var s string
	//var s2 string
	//
	//// Send in the reference of the string s
	//fmt.Scanln(&s, &s2)
	//// Print the value of it.
	//fmt.Println(s, s2)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	// For a float number
	fmt.Print("Enter number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
}
