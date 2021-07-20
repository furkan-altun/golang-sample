package main

import "fmt"

func main() {
	fmt.Println("Number\tHexa   Value")
	fmt.Println("------\t------ -----")
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t\t%#U\n", i, i)
	}
}
