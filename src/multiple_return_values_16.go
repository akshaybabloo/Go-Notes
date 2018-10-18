package main

import "fmt"

func main() {

	n1, l1 := FullName("akshay", "gollahalli")
	fmt.Printf("Fullname: %v, number of characters: %v\n", n1, l1)

	n2, l2 := FullNameNakedReturn("akshay", "gollahalli")
	fmt.Printf("Fullname: %v, number of characters: %v\n", n2, l2)

}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)

	return full, length
}

func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)

	return
}