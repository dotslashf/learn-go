package main

import "fmt"

func arraySlicesMap() {
	var listOfNames [4]string
	listOfNames[0] = "John"
	listOfNames[1] = "Doe"
	listOfNames[2] = "Jane"
	for i := 0; i < len(listOfNames); i++ {
		fmt.Println("Name index:", i, "is", listOfNames[i])
	}

	var listOfFruits = [4]string{
		"Apple",
		"Banana",
		"Watermelon",
		"Orange",
	}
	for i := 0; i < len(listOfFruits); i++ {
		fmt.Println("Fruits index:", i, "is", listOfFruits[i])
	}
	fmt.Println("Fruits:", listOfFruits)

	var numbers = [...]int8{6, 2, 0, 4, 2, 0}
	fmt.Println("Numbers:", numbers)
	fmt.Println("Length of numbers:", len(numbers))

	var numbersMultidimensional = [2][5]int8{{4, 1, 3}, {4, 5, 6, 7, 8}}
	for i := 0; i < len(numbersMultidimensional); i++ {
		for j := 0; j < len(numbersMultidimensional[i]); j++ {
			if numbersMultidimensional[i][j] != 0 {
				fmt.Printf("Current value %d ", numbersMultidimensional[i][j])
			}
		}
		fmt.Println("")
	}

	for i, fruit := range listOfFruits {
		fmt.Printf("Index %d, Fruits %s\n", i, fruit)
	}

	var emojis = make([]string, 3)
	emojis[0] = "🤓"
	emojis[1] = "💀"
	emojis[2] = "🔥"

	for _, emoji := range emojis {
		fmt.Printf("Emoji %s\n", emoji)
	}

}
