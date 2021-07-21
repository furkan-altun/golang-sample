package main

import "fmt"

func main() {
	i := []int{11, 22, 33, 44, 55, 66, 77, 88, 999, 10}
	for k, v := range i{
		fmt.Println(k+1,"th value is:", v)
	}
	fmt.Printf("%T", i)
}
