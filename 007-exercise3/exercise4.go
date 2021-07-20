package main

import (
	"fmt"
	"time"
)

func main() {
	myBirthdayYear := 1994
	currentYear := time.Now().Year()
	for {
		// bu şekilde yapılan for loop syntax i içinde break konulmazsa sonsuz döngüye giriyor.
		// o sebeple çıkış case i için kontrol koyulup, o case e girdiğinde break diyip çıkılmalı.
		if myBirthdayYear > currentYear {
			break
		}
		fmt.Println(myBirthdayYear)
		myBirthdayYear++
	}
	fmt.Println("--------------")
}
