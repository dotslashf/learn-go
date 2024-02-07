package main

import "fmt"

func constant() {
	const firstName string = "fadhlu"
	const lastName = "rahman"
	
	// multiple constant with type inference
	const (
		country = "Indonesia"
		isRaining = false
		numberInt = 1
		numberFloat = 3.142
	)

	// multiple constant type inference from prev variable
	const (
		today string = "Wednesday"
		tomorrow
		isNewYear = false
	)

	const numberOne, numberTwo = 1, 2
	const planet, star = "Earth", "Sun"

	fmt.Println("My name is", firstName, lastName)
	fmt.Println(country, "is my country")
	fmt.Println("Is today raining?", isRaining)
	fmt.Println("How many is god?", numberInt)
	fmt.Printf("Value of pi? %f", numberFloat)
}