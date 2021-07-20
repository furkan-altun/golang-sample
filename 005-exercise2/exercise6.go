package main

import (
	"fmt"
)

// var currentYear int = time.Now().Year()

const (
	// e6 = iota + currentYear
	e6 = iota + 2021
	e7
	e8
	e9
)

func main() {
	fmt.Println(e6)
	fmt.Println(e7)
	fmt.Println(e8)
	fmt.Println(e9)
}
