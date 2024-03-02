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

	var numbersInt = [...]int8{6, 2, 0, 4, 2, 0}
	fmt.Println("Numbers:", numbersInt)
	fmt.Println("Length of numbers:", len(numbersInt))

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

	var languages = []string{"Javascript", "Java", "Go"}
	fmt.Println(languages[0])

	var languagesSlice = []string{"Javascript", "Java"}  //slice
	var languagesArray = [2]string{"Go", "Rust"}         //array
	var languagesArrayDot = [...]string{"Elixir", "C++"} // array
	fmt.Println(languagesSlice, languagesArray, languagesArrayDot)

	var nationalities = []string{"Indonesian", "Singaporean", "South Korean"}
	var seaNationalities = nationalities[0:2]
	fmt.Println(seaNationalities)

	var numbers = []string{"0", "1", "2", "3"}
	var numbersA = numbers[0:3]
	var numbersB = numbers[1:4]

	var numbersAA = numbersA[1:2]
	var numbersAB = numbersB[0:1]

	fmt.Println(numbers, numbersA, numbersB)
	fmt.Println(numbersAA, numbersAB)

	numbersAB[0] = "69"

	fmt.Println(numbers, numbersA, numbersB)
	fmt.Println(numbersAA, numbersAB)

	var carBrands = []string{"Honda", "Toyota", "Wuling", "Mitsubishi"}
	fmt.Println(len(carBrands), cap(carBrands))

	var sliceCarBrands = carBrands[0:3]
	fmt.Println(len(sliceCarBrands), cap(carBrands))

	var anotherSliceCarBrands = carBrands[1:4]
	fmt.Println(len(anotherSliceCarBrands), cap(anotherSliceCarBrands))

	var newCarBrands = append(carBrands, "Daihatsu")
	fmt.Println(newCarBrands)

	var twoCarBrands = carBrands[0:2]
	fmt.Println(len(twoCarBrands), cap(twoCarBrands))
	fmt.Println(twoCarBrands, carBrands)

	var addedCarBrands = append(twoCarBrands, "Suzuki")
	fmt.Println(carBrands, twoCarBrands, addedCarBrands)

	var dstCars = make([]string, 3)
	var src = carBrands[0:3]
	var cpyCars = copy(dstCars, src)
	fmt.Println(dstCars, cpyCars)

	var blankCars = []string{"blank", "blank", "blank"}
	var srcCars = carBrands[0:2]
	copy(blankCars, srcCars)
	fmt.Println(blankCars, srcCars)

	var colors = []string{"Red", "Blue", "Green"}
	var colorsA = colors[0:2]
	var colorsB = colors[0:2:2] // index start : index end : cap

	fmt.Println(colors, len(colors), cap(colors))
	fmt.Println(colorsA, len(colorsA), cap(colorsA))
	fmt.Println(colorsB, len(colorsB), cap(colorsB))

}
