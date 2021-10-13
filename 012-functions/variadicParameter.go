package main

import "fmt"

func main()  {
	sumOfMyNumbers := sum(1, 3, 5, 7, 9)
	fmt.Println("Sum of my numbers are : ", sumOfMyNumbers)
}

func sum(numbers ...int) int {
	sum := 0
	for index, value := range numbers {
		sum += value
		fmt.Printf("%d. index of numbers is: %d and sum of numbers are %d now\n", index, value, sum)
	}
	return sum
}
