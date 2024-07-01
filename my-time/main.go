package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time!")

	// Getting the current time
	presentTime := time.Now()
	fmt.Println(presentTime)

	// formatting time
	// reference time Mon Jan 2 15:04:05 MST 2006
	// 02 - day
	// 01 - month
	// 2006 - year
	fmt.Println(presentTime.Format("02-01-2006 Monday"))

	// Working with slices
	arr := []string{"wild", "solid", "sauvage", "spirit", "soul"}
	fmt.Println("Welcome to slices!", arr)

	// Removing the third element ("sauvage")
	arrN := append(arr[:2], arr[3:]...)
	fmt.Println("Updated slice:", arrN)

	// Working with maps
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("Ages map:", ages)

	// Iterating over the map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Current Unix time in nanoseconds
	now := time.Now()
	nanos := now.UnixNano()
	fmt.Println("Current Unix time in nanoseconds:", nanos)

}