package main

import "fmt"

type vehicle struct {
	door string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck1 := truck{
		vehicle : vehicle{
			door: "sample-door",
			color: "black",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle : vehicle{
			door: "sample-door",
			color: "red",
		},
		luxury: false,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1.vehicle.color, sedan1.vehicle.door, sedan1.door, sedan1.luxury)

	// Anonymous struct
	bird1 := struct{
		age int
		color string
		carnivorous bool
		qualifications []string
	}{
		age: 10,
		color: "gray",
		carnivorous: true,
		qualifications: []string {"run", "fly", "swim"},
	}

	fmt.Println(bird1.age, bird1.color, bird1.carnivorous)
	for _, v := range bird1.qualifications{
		fmt.Printf("%s | ", v)
	}
}
