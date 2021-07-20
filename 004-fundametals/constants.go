package main

import "fmt"

const a int = 42
const b string = "Furkan"
const c float32 = 58.423

const (
	e int = 53
	r string = "Furkan Altun"
	t float64 = 534.321
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", r)
	fmt.Printf("%T \n", t)
}
