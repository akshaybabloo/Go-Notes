package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "./hello.txt"

	content, err := ioutil.ReadFile(fileName) // returns bytes
	checkError(err)
	byteToString := string(content)
	fmt.Println("Read from file", byteToString)
}


func checkError(err error) {
	if err != nil {
		panic(err)
	}
}