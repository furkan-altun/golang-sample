package main

import "fmt"

const i = 7.5
const j = 42
const k = "A"

func main() {
	fmt.Printf("decimal: %f \n", i)
	fmt.Printf("binary: %b \n", j)
	fmt.Printf("hexa: %#x \n", k)
	fmt.Printf("hexa: %#x \n", j)
	fmt.Printf("%T \n", i)
	fmt.Printf("%T \n", j)
	fmt.Printf("%T \n", k)
}
