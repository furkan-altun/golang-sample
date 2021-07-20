package main

import (
	"fmt"
	"time"
)

func main() {
	myBirthdayYear := 1994
	currentYear := time.Now().Year()
	for myBirthdayYear <= currentYear {
		fmt.Println(myBirthdayYear)
		myBirthdayYear++
	}
}
