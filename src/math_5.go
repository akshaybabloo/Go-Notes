package main

import (
	"fmt"
	"math/big"
	"math"
)

func main() {
	int1, int2, int3 := 10, 20, 30
	intSum := int1 + int2 + int3
	fmt.Println("Int sum:", intSum)

	fl1, fl2, fl3 := 10.1, 20.2, 30.3
	floatSum := fl1 + fl2 + fl3
	fmt.Println("Float sum:", floatSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(10.1)
	b2.SetFloat64(20.2)
	b3.SetFloat64(30.3)
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("Big sum: %.10g\n", &bigSum)

	circleRadius := 20.1
	circumference := circleRadius * math.Pi
	fmt.Printf("circumference %.2f\n", circumference)
}
