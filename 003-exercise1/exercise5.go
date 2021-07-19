package main

import "fmt"

type hotdog int

var x5 hotdog
var y5 int

func main() {
	fmt.Println(x5)
	fmt.Printf("%T \n", x5)
	x5 = 58
	fmt.Println(x5)

	y5 = int(x5)
	fmt.Println(y5)
}
