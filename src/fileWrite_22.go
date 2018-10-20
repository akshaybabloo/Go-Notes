package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is a string"

	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("File written with %v characters", ln)

	bytes := []byte(content)
	err1 := ioutil.WriteFile("./fromBytes.txt", bytes, 0644) // 0644 is the file permission
	checkError(err1)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
