package main

import "fmt"

func main() {
	// Array in kapasitesi belli ise array, değil ise Slice deniyor.
	var x [5]int;
	fmt.Println(x)
	x[0] = 4
	x[3] = 7
	fmt.Println(x)

	fmt.Println("------------------------")

	//SLICES
	y := []int{5,3,7,8,2}
	fmt.Println(y)
	fmt.Println(len(y))

	fmt.Println("------------------------")

	for i, v := range y {
		fmt.Println(i, v)
	}

	fmt.Println("------------------------")

	for _, v := range y {
		fmt.Println(v)
	}

	fmt.Println("------------------------")
}
