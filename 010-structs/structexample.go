package main

import "fmt"

type person struct {
	firstname string
	lastname string
	age int
}

// Embedded structs
type secretAgent struct {
	//person --> eğer field ismi belirtilmemişse sadece type ismi belirtilmişse, field ismi otomatik olarak type ismi olur.
	fbiPerson person
	licenceToKill bool
}

func main() {
	p1 := person{
		firstname: "Furkan",
		lastname: "Altun",
		age: 27,
	}

	p2 := person{
		firstname: "Valentin",
		lastname: "Rosier",
		age: 24,
	}

	secretAgent1 := secretAgent{
		fbiPerson: person{
			firstname: "James",
			lastname: "Bond",
			age: 32,
		},
		licenceToKill: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(secretAgent1)

	fmt.Println(p1.firstname, p1.lastname, p1.age)
	fmt.Println(p2.firstname, p2.lastname, p2.age)
	fmt.Println(secretAgent1.fbiPerson.firstname, secretAgent1.fbiPerson.lastname, secretAgent1.fbiPerson.age, secretAgent1.licenceToKill)

	// Anonymous Struct
	// Anonymous struct function içinde composite literal olarak tanımlanır. Func dışında yeni bir type oluşturularak
	// tanımlanmaz. Kullanım alanı çok az olacak struct tanımlamaları için ayrı ayrı typelar oluşturmak yerine
	// bu şekilde hızlıca tanım yaparak daha clean bir code yazılması amaçlanır.
	a1 := struct {
		color string
		shape string
		length int
		height int
	}{
		color: "black",
		shape: "rectangle",
		length: 100,
		height: 75,
	}
	fmt.Println(a1)
}
