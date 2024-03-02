package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func function() {
	var names = []string{"John", "Jane"}
	printNames("Hello", names)
	fmt.Println("Random number 1:", randomWithRange(1, 100))
	fmt.Println("Random number 2:", randomWithRange(1, 10))
	fmt.Println("Random number 3:", randomWithRange(1, 5))
	fmt.Printf("Divide result of %d and %d is %d\n", 6, 2, divide(6, 2))
}

func printNames(message string, names []string) {
	fmt.Println(message, strings.Join(names, ", "))
}

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func randomWithRange(min int, max int) int {
	return randomizer.Int()%(max-min+1) + min
}

func divide(a, b int) int {
	if b == 0 {
		fmt.Println("Invalid divider, cannot divide by zero")
	}
	return a / b
}
