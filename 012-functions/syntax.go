package main

import (
	"fmt"
	"time"
)

func main()  {
	sayMyName()
	saySurname("ALTUN")
	fa := getFullNameAndAge()
	fmt.Println(fa)
	fullName, birthOfYear := getPersonelInfo("Valentin","Rosier",24)
	fmt.Println(fullName, birthOfYear)
}

func sayMyName()  {
	fmt.Println("Furkan")
}

func saySurname(surname string)  {
	fmt.Printf("Surname\t: %s\n",surname)
}

func getFullNameAndAge() string {
	return fmt.Sprintf("Name\t:%s\nSurname\t:%s\nAge\t:%d","Ricardo","Quaresma",18)
}

func getPersonelInfo (firstName string, lastName string, age int) (string, int)  {
	fullName := fmt.Sprintf("%s - %s \n", firstName, lastName)
	yearOfBirth := time.Now().Year() - age
	return fullName, yearOfBirth
}