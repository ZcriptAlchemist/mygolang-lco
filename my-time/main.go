package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time!")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// formatting time
	// reference time Mon Jan 2 15:04:05 MST 2006
	// 02 - day
	// 01 - month
	// 2006 - year
	fmt.Println(presentTime.Format("02-01-2006 Monday"))
}