// strings and runes
// runes are also int32, but is better to be explicit to to clarify intent

// Go does not allow truthiness, to convert to another type to boolean, use comparison operators.

// :=
// // is not legal outside of functions
// // Empty str or 0 int, use var
// // Let assign new variables or rehuse existing ones
// // You can use type a conversion, but it is more idimatic to use var

// Avoid declaring functions outside of functions

// const is for inmutable values

package main

import "fmt"

func main() {

	// Another way to declare multiple const or var
	const (
		owner = "Gustavo"
		age   = 21
	)
	var (
		langs int    = 2
		food  string = "chilaquiles"
	)

	// Using :=
	city, cp := "Mexico", 55341

	fmt.Printf("My name is %v, i am %v yo\n", owner, age)
	fmt.Printf("I speak %v languages and my fav food are %v\n", langs, food)
	fmt.Printf("I live in %v and the cp is %v\n", city, cp)

	city, job := "Canada", "SW" // Rehusing city, this sintax only if a new var is created
	fmt.Printf("I want to live in %v working in %v\n", city, job)

}
