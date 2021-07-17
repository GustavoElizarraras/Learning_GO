package main

import "fmt"

func main() {

	// for is the only looping keyword in the language

	// Complete for statement

	// The comparison occurs before the loop begin
	for i := 0; i < 10; i++ { // No parenthesis, must use := to initialize the variables
		// var is not legal
		fmt.Println(i)
	}

	// Condition-Only for statement, simulate a while in other languages
	j := 1
	for j < 100 {
		fmt.Println(j)
		j *= 2 // Same as other languages, operator and equal
	}

	// Infinite for
	// for {
	// 	fmt.Println("Hello")
	// }

	// break, the loop is exited
	// continue, the loop ignore what is next in that iteration, goes to the next.

	// for-range statement, iterate over built-in types, it is a copy

	primeNumbers := []int{3, 7, 11, 17, 23}
	for i, v := range primeNumbers {
		fmt.Println(i, v) // Two varaibles using range
	}
	fmt.Println()

	nums := map[string][]int{
		"one":   {1},
		"two":   {2},
		"three": {3},
		"four":  {4},
	}

	// Nested for loops

	for k, v := range nums { // to ignore the key or index, use "_"

		fmt.Printf("Key: %v and Value: %v\n", k, v)

		// to ignore the value, only one varaible
		for k := range nums {
			fmt.Printf("Key: %v\n", k)
		}
	}
	fmt.Println()
	// The order of Hash maps are not fixed when iterating through them. So
	// a Hash DoS attack is harder to implement

	// Iterating over strings

	var s string = "apple_π!"
	// strings are iterated over the runes, not the bytes
	// the pi string is converted to a 32-bit number
	for i, run := range s {
		fmt.Printf("%v, %v, %v\n", i, run, string(run))
	}

	// Labeling, they are in the same indent level as the block that contains them

	sl_strs := []string{"Hello ", "apple_π!"}
	// iterating over a slice, then over each string
outer:
	for _, word := range sl_strs {
		for i, run := range word {
			fmt.Println(i, run, string(run))
			if run == 'l' {
				continue outer // It goes to the next word if the rune is an l
			}
		}
	}

}
