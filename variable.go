package main

import "fmt"

func variable() {
	var firstName string = "fadhlu"
	var lastName string = "rahman"
	fullName := firstName + lastName // variable with type inference
	fullName = "ini budi" // reassign value no need using :
	var longMessage string = `
This a message with new line
Its so long that it used backtick
`

	var one, two, three int = 1, 2, 3
	four, five, six := 4, 5, 6

	var positiveNumber uint8 = 77
	negativeNumber := -69
	var negativeFloatNumber float32 = -69.420
	var isHuman bool = true

	/*
	multiple variable with type inference

	numberOne, isBoolean, numberFloat, words := 1, true, 3.14, "ayyo"
	*/

	fmt.Printf("Halo saya %s%s\n", firstName, lastName) // format print
	fmt.Println("Halo saya", firstName + lastName) // normal print
	fmt.Printf("Nama lengkap saya adalah %s\n", fullName)
	fmt.Printf("This is number %d, %d, %d, %d, %d, %d\n", one, two, three, four, five, six)
	fmt.Printf("Positive number %d\n", positiveNumber)
	fmt.Printf("Negative number %d\n", negativeNumber)
	fmt.Printf("Negative float number %.3f\n", negativeFloatNumber)
	fmt.Printf("Is %s a human? %t", fullName, isHuman)
	fmt.Println(longMessage)
}