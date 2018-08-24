package main

import (
	"time"
	"fmt"
)

func main() {
	// Format date
	t := time.Date(2010, time.November, 12, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Date and time %s\n", t)
	// Splitting t
	fmt.Println("Year", t.Year())
	fmt.Println("Month", t.Month())
	fmt.Println("Day", t.Day())
	fmt.Println("Week", t.Weekday())

	// Current date and time
	currentDateTime := time.Now()
	fmt.Printf("Date and time now %s\n", currentDateTime)

	tomorrow := currentDateTime.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n",
		tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longFormat := "Monday, January, 2, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

	shortFormat := "2/1/06"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))
	
}
