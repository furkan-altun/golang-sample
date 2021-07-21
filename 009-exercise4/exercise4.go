package main

import "fmt"

func main() {
	i4 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	i4 = append(i4, 52)
	fmt.Println(i4)
	i4 = append(i4, 53, 54, 55)
	fmt.Println(i4)
	y4 := []int{56, 57, 58, 59, 60}
	i4 = append(i4, y4...)
	fmt.Println(i4)
}
