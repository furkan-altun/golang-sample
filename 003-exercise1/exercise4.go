package main

import "fmt"

type furkan int
var x furkan

func main() {
	fmt.Println(x)
	fmt.Printf("%T \n", x)
	x = 42
	fmt.Println(x)
}
