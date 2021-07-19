package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Printf("%v", x) // the value in a default format when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%t", x) // %t the word true or false
}
