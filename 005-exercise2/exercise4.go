package main

import "fmt"

func main() {
	a4 := 42
	fmt.Printf("%d %b %#x \n", a4, a4, a4) // 42 --> 101010
	b4 := a4 << 1
	fmt.Printf("%d %b %#x \n", b4, b4, b4) // 42 nin binaryde bir adım sola kaydırılmış hali 1010100
	c4 := b4 << 1
	fmt.Printf("%d %b %#x \n", c4, c4, c4) // 42 nin binaryde iki adım sola kaydırılmış hali 10101000
}

// 2^0 = 1 0  -->
// 2^1 = 2 10 --> 1 << x
// 2^2 = 4 100 --> 1 << y
// 2^3 = 8 1000 --> 1 << z