package main

import "fmt"

func main() {
	var firstName string = "fadhlu"
	var lastName string = "rahman"
	fullName := firstName + lastName // variable with type inference
	fullName = "ini budi" // reassign value no need using :
	
	var one, two, three int = 1, 2, 3
	four, five, six := 4, 5, 6

	/*
	multiple variable with type inference
	
	numberOne, isBoolean, numberFloat, words := 1, true, 3.14, "ayyo"
	*/

	fmt.Printf("Halo saya %s%s\n", firstName, lastName) // format print
	fmt.Println("Halo saya", firstName + lastName) // normal print
	fmt.Printf("Nama lengkap saya adalah %s\n", fullName)
	fmt.Printf("This is number %d, %d, %d, %d, %d, %d", one, two, three, four, five, six)
}