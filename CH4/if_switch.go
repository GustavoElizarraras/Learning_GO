package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// If statements are like other programming languagesconst
	for i := 0; i < 5; i++ { // complete for
		n := rand.Intn(10)    // first always set to one
		if n == 1 && i == 0 { // no parenthesis again
			fmt.Println("It is hard coded")
		} else if n > 5 {
			fmt.Println(n, " is bigger than 5")
		} else {
			fmt.Println(n, " is less than 6")
		}
	}

	// switch
	vehicle := map[string][]string{
		"car":   {"earth"},
		"bike":  {"earth"},
		"boat":  {"water"},
		"plane": {"air"},
		"piano": {"music"},
	}

	for k, w := range vehicle {
		// fmt.Println(k, w[0])
		switch w[0] {
		case "earth":
			fmt.Println(k, " in ", w)
		case "water":
			fmt.Println(k, " in ", w)
		case "air":
			fmt.Println(k, " in ", w)
		default:
			fmt.Println(k, " is not a vehicle")
		}
	}

	words := []string{"a", "cow", "smile", "octopus", "paleonthologist"}
	for _, word := range words {
		switch size := len(word); size { // size is a variable with scope inside the switch
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short one")
		case 5:
			fmt.Println(word, " is a 5 letter word")
		case 6, 7, 8, 9: // empty case, do nothing
		default:
			fmt.Println(word, " is a long word")
		}

		// break can be used to exit early from a case, but that need indicate something
		// difficult and may not be the best option, refactor

		// Breaking an outer for may be a requierment
	loop:
		for i := 1; i < 22; i *= 2 {
			switch { // black switch, boolean comparison for each case
			case i == 1:
				fmt.Println("one")
			case i == 8:
				fmt.Println("exiting for")
				break loop
			default:
				fmt.Println("other")
			}
		}
		// If comparing with the same variable in all cases, is better a switch statement

		// goto also exist, it is recommended not to use it, but there are special cases where
		// it is useful, like avoiding duplicate complicated code.
	}

}
