package main

import (
	"fmt"
	"stringutil"
)

func main() {

	n1, l1 := stringutil.FullName("akshay", "gollahalli")
	fmt.Printf("Fullname: %v, number of characters: %v\n", n1, l1)

	n2, l2 := stringutil.FullNameNakedReturn("akshay", "gollahalli")
	fmt.Printf("Fullname: %v, number of characters: %v\n", n2, l2)

}
