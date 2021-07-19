package main

import (
	"fmt"
	"sort"
)

// Data validation and error checking is important

// return Functions from functions (return a closure)
func makeMul(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	// Anonymus functions, a function inside a function, they don not have a name
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "inside of an anonymous function")
		}(i) // calling the function with parenthesis
		// These are mostly used with defer and launching goroutines
	}
	// Closures
	// Functions declared inside functions that are able to access and modify variables
	// declared in the outer function. Useful when a function call another multiple times,
	// the inner function can be used to "hide" the called function.

	// Passing functions as parameters
	// Treat functions like data, a created closure that references a local variable and then
	// passing that closure to another functions.+

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Peterson", 37},
		{"Ted", "Bobbert", 72},
		{"Marshall", "B.Fudge", 21},
	}
	fmt.Println(people)
	// Sorting slice by last name
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	// Closure that is passed to sort.Slice has 2 parameters, i and j, but within
	// we refer to people so we can sort it by the LastName field. People is "captured" by the closure.
	fmt.Println(people)

	// Sorting slice by age
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
	// Passing functions as parameters to other funcs is useful for performing different operations
	// on the same kind of data

	r1, r2 := makeMul(2), makeMul(4)
	for i := 0; i < 4; i++ {
		fmt.Println("Return a function from a function: ", r1(i), r2(i))

	}

}
