package main

import "fmt"

func main() {
	var i [5]int
	i[0] = 77
	i[1] = 33
	i[2] = 22
	i[3] = 11
	i[4] = 99

	fmt.Println(i)

	for k, v := range i{
		fmt.Println(k, v)
	}

	fmt.Printf("%T", i)
}
