package main

import "fmt"

func main() {
	k1 := []string {"James", "Bond", "Shaken, not stirred"}
	k2 := []string {"Miss", "Moneypenny", "Helloooooo, James."}

	k7 := [][]string{k1, k2}
	fmt.Println(k7)

	for k, v := range k7 {
		fmt.Println(k, v)
		for k2, v2 := range v {
			fmt.Println(k2, v2)
		}
	}
}
