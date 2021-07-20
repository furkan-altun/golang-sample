package main

import (
	"fmt"
	"time"
)

func main() {
	currentYear := time.Now().Year()

	if currentYear > 2000 {
		fmt.Println("You are in after 2000")
	}else if currentYear == 2000{
		fmt.Println("You are in 2000")
	}else {
		fmt.Println("You are in before 2000")
	}
}
