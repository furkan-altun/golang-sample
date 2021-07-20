package main

import "fmt"

func main() {
	i := "Orta Saha"
	switch i {
	case "Kaleci":
		fmt.Println("Ersin")
	case "Defans":
		fmt.Println("Rosier")
	case "OrtaSaha":
		fmt.Println("Atiba")
	case "Forvet":
		fmt.Println("Larin")
	default:
	 	fmt.Println("Sergen Yalçın")
	}

	// default case and it may appear anywhere in the "switch" statement
	switch {
	case true:
		fmt.Println("It should print")
	case false:
		fmt.Println("It shouldnt print")
	}
}
