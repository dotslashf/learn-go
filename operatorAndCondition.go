package main

import "fmt"

func operatorAndCondition() {
	const value int8 = (((2 + 6) % 3) * 4 - 2 ) / 3
	const isEqual bool = value == 2

	const left bool = true
	const right bool = false

	leftAndRight := left && right
	leftOrRight := left || right
	notLeft := !left

	fmt.Printf("Value of %d (%t)\n", value, isEqual)
	fmt.Printf("Value of leftAndRight (%t), leftOrRight (%t), notLeft (%t)\n", leftAndRight, leftOrRight, notLeft)

	const finalValue int8 = 10
	
	if finalValue == 10 {
		fmt.Println("Perfect Score")
	} else if finalValue >= 5 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Take retest")
	}

	const score float32 = 8884.0

	if percentage := score / 100; percentage >= 100 {
		fmt.Printf("Perfect, %1.f%s\n", percentage, "%")
	} else if percentage >= 50 {
		fmt.Printf("Passed, %1.f%s\n", percentage, "%")
	} else {
		fmt.Printf("Not Passed, %1.f%s\n", percentage, "%")
	}

	const point int8 = 79

	switch {
	case point == 100:
		fmt.Println("Score: A")
	case (point >= 80) && (point < 100):
		fmt.Println("Score: B")
	case (point >= 60) && (point < 80):
		fmt.Println("Score: C")
		fallthrough
	default:
		fmt.Println("Score: E")
	}
}