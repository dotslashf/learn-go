package main

import "fmt"

func looping() {
	for i := 0; i <= 5; i++ {
		fmt.Println("Current i", i)
	}

	var i int8 = 0
	for i < 5 {
		fmt.Println("Current ii", i)
		i++
	}

	i = 0
	for {
		fmt.Println("For current i without argument", i)

		i++
		if i == 5 {
			break
		}
	}

	var chars string = "abc123456"
	for i, v := range chars {
		fmt.Printf("Value %d, Index %d\n", v, i)
	}

	var integers = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range integers {
		fmt.Printf("Value %d\n", v)
	}

	var slices = integers[0:2] // slices
	for i, v := range slices {
		fmt.Printf("Value %d, Index %d\n", v, i)
	}

	// hashes = map[datatype for key]value
	var hashes = map[string]int{"a": 4, "b": 5, "c": 6} // hashmap
	for k, v := range hashes {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	for range hashes {
		fmt.Println("Done")
	}

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue // line after continue is ignored
		}

		if i > 8 {
			break
		}

		fmt.Printf("Value of i: %d\n", i)
	}

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

outerloop:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 3 {
				break outerloop
			}
			fmt.Printf("matrix [%d][%d]\n", i, j)
		}
	}
}
