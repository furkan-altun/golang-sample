package main

import "fmt"

func main() {
	i8 := map[string] []string{
		"altun_furkan" : {"football", "computer games", "learning programming"},
		"rosier_valentin" : {"playing music", "singing", "playing football" },
		"hutchinson_atiba" : {"running", "yoga", "swimming"},
	}

	//add record to map
	i8["larin_cycle"] = []string {"test", "play guitar", "playing football"}

	//delete record from map
	delete(i8, "altun_furkan")

	for k, list := range i8 {
		fmt.Print(k, ": ")
		for _,v := range list {
			fmt.Print(v, " | ")
		}
		fmt.Print("\n")
	}

}
