package main

import "fmt"

func main() {
	s := "Hello, my name is Furkan"
	fmt.Println(s)
	s = "Furkan Altun"
	fmt.Printf("%v \n", s)
	fmt.Printf("%T \n",s)

	for ind, val := range s {
		fmt.Println(ind, val)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}
}