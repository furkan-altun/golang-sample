package main

import "fmt"

func main() {
	for i := 10; i <= 100 ; i++ {
		fmt.Printf("The number is: %d | remainder diveded by four is: %d \n", i,  i % 4)
	}
}
