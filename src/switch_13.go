package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	//dow := rand.Intn(6) + 1
	//fmt.Println("Day", dow)

	result := ""

	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "It's Sunday!"
	case 2:
		result = "It's Monday!"
	case 3:
		result = "It's Tuesday!"
	case 4:
		result = "It's Wednesday!"
	case 5:
		result = "It's Thursday!"
	case 6:
		result = "It's Friday!"
	default:
		result = "It's Saturday!"
	}

	//fmt.Println("Day", dow, result)
	fmt.Println(result)

	// Or you can completely eliminate the switch expression
	x := 30

	switch {
	case x < 0:
		result = "Less than 0"
	case x > 0:
		result = "Greater than 0"
		//fallthrough // executes this code and also the next one.
	default:
		result = "Equal to 0"
	}

	fmt.Println(result)
}
