package main

import "fmt"

func main() {
	fmt.Println("Hello, world. It is beginnig...")
	printEvenNumbers()
	fmt.Println("End of the code!")
}

func printEvenNumbers ()  {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}
