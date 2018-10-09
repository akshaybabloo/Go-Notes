package main

import (
	"fmt"
	"sort"
)

func main() {
	countries := make(map[string]string)

	countries["NZ"] ="New Zealand"
	countries["IN"] ="India"
	countries["AU"] ="Australia"
	countries["USA"] ="United States of America"
	countries["UK"] ="United Kingdom"

	fmt.Println(countries)
	fmt.Println(countries["NZ"]) // Retrieve a value

	delete(countries, "AU") // Delete a key and its value
	fmt.Println(countries)

	// Looping over maps
	for k,v := range countries {
		fmt.Printf("%v: %v\n", k, v)
	}

	// Alphabetical order
	keys := make([]string, len(countries))
	i := 0
	for k:= range countries {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	// Sort the keys
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys{
		fmt.Println(countries[keys[i]])
	}
}
