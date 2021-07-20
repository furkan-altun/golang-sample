package main

import "fmt"

const (
	x = 10
	y = 15
	z = 10
)

func main() {
	a1 := x == z // true
	a2 := x != y // true
	a3 := y <= z // false
	a4 := x >= z // true
	a5 := y < z // false
	a6 := y > x // true

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a6)
}