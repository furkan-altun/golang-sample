package main

import "fmt"

var y string
var z int

//C# default value = Go Zero Value
// false for booleans,
// 0 for integers,
// 0.0 for floats,
// "" for strings,
// nil for pointers, functions, interfaces, slices, channels, and maps.
func main() {
	fmt.Println("y nin değeri: -", y, "-")
	fmt.Printf("y nin tipi: %T \n", y)

	y = "Furkan Altun, Beşiktaş"

	fmt.Println(y)
	fmt.Printf("%T \n", y)

	fmt.Println("z nin değeri: ", z)
	fmt.Printf("z nin tipi: %T \n", z)
}
