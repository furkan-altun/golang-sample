package main

import "fmt"

func main() {
	switch {
		case true:
			fmt.Println("This should print")
		case false:
			fmt.Println("This shouldn't print")
	}

	k7 := 66
	switch {
		case k7 < 18:
			fmt.Println("You are teenager")
		case k7 > 18 && k7 < 60 :
			fmt.Println("You are middle-aged")
		default:
			fmt.Println("You are older")
	}

	k8 := 19
	switch k8 {
		case 18:
			fmt.Println("You are teenager")
		case 35:
			fmt.Println("Middle of life")
		case 70:
			fmt.Println("End of life")
		default:
			fmt.Println("Undescribed age")
	}
}
