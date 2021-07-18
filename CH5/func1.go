package main

import (
	"errors"
	"fmt"
)

// Simple function
func div1(num int, den int) int { // func, name(args with type) return and it's type
	// If a function returns a value (in declaration), you must supply a return
	// If it does not return a value, left blank between the args and braces
	if den == 0 {
		return 0
	}
	return num / den
}

func div2(num, den int) int { // If the args are of the same type, they can be declared like this
	if den == 0 {
		return 0
	}
	return num / den
}

// Variadic Input Parameters and Slices
// The variadic parameter must be the last or only arg in the func, 3 dots before the type: ...type
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals)) // slice int variable, length depends from the second to last number (or slice) provided
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// Simulating Named and Optional Parameters
// Go does not have this feature, but it can be simulated with a struct
// If a function has too many parameters, it may be too complicated

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

// func Myfunc(opts MyFuncOpts) error {
// do something here
// Inside main it should be called like this:
// // 	Myfunc(MyFuncOpts{
// // 		LastName: "Patel",
// // 		Age:      20,
// // 	})

// // 	Myfunc(MyFuncOpts{
// // 		FirstName: "JJ",
// // 		LastName:  "Patel",
// // 	})
// }

// Multiple Return Values are Multiple Values, explicit
func divAndRemainder(num int, den int) (int, int, error) { // returns are listed in parentheses
	if den == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / den, num % den, nil // here the parenthesis are not needed, compile time error
	// all must be returned, separated by commas
	// if a return variable is not needed, when calling the func, use "_"
}

func main() {

	// Calling div and assign the result to a varaible
	r1 := div1(5, 2)
	r2 := div2(10, 2)
	fmt.Println("div1, 5 / 2 is : ", r1)
	fmt.Println("div 2, 10 / 2 is : ", r2)
	// Calling addTo()
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // 3 dots needed, if not, compile-time error
	// Calling multiple return values func
	r3, remain3, err := divAndRemainder(9, 3)
	fmt.Println(r3, remain3, err)
}
