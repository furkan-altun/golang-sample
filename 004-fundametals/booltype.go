package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Printf("%v \n", x) // the value in a default format when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%t \n", x) // %t the word true or false

	a := 7
	b := 10

	fmt.Println(a == b) // return false
	fmt.Println(a != b) // return true
}
