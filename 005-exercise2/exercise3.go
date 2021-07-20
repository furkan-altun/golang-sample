package main

import "fmt"

const (
	untypedConst1 = 2
	untypedConst2 = "Furkan"
	untypedConst3 = true
)

const (
	typedConst1 int = 3
	typedConst2 string = "Altun"
	typedConst3 bool = false
)

func main() {
	fmt.Printf("Untyped Constants:\t%v - %T | %v - %T | %v - %T \n",
		untypedConst1, untypedConst1, untypedConst2, untypedConst2, untypedConst3, untypedConst3)
	fmt.Printf("Typed Constants:\t%v - %T | %v - %T | %v - %T \n",
		typedConst1, typedConst1, typedConst2, typedConst2, typedConst3, typedConst3)
}