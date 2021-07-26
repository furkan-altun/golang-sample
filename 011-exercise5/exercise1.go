package main

import "fmt"

type person struct {
	firstName string
	lastName string
	favoriteIceCreamFlavor []string
}

func main() {
	p1 := person{
		firstName: "Furkan",
		lastName: "Altun",
		favoriteIceCreamFlavor: []string{"Melon", "Lemon"},
	}

	p2 := person{
		"Valentin",
		"Rosier",
		[]string{"Chocolate","Vanilla"},
	}

	fmt.Println(p1)
	fmt.Println(p2.firstName, p2.lastName, p2.favoriteIceCreamFlavor)

	m := map[string] person{
		p1.lastName : p1,
		p2.lastName : p2,
	}

	for i, v := range m{
		fmt.Println(i, ": ", v.firstName, v.lastName, v.favoriteIceCreamFlavor)
	}

}
