package main

import "fmt"

func main() {
	//region if statements
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	if 5>0 {
		fmt.Println("005")
	}
	if 0<1 {
		fmt.Println("006")
	}
	if a := 1; a < 2 {
		fmt.Println("007")
	}
	//endregion if statements

	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}
