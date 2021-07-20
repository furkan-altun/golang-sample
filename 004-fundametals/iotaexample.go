// You could do some interesting things with iota if you need something to automatically increment by
package main

import "fmt"

const (
	a1 = iota
	a2
	a3
)

const (
	b1 = iota
	b2
	b3
	b4 = iota
	b5
	b6
)

func main() {
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
	fmt.Println(b5)
	fmt.Println(b6)
}
