package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Named Return Values

func divAndRemainder(num int, den int) (res int, rem int, err error) { // parentheses, even if it is one
	// pre-declaring variables that you use within the function to hold the return values
	// they are initialized with their zero value
	if den == 0 {
		err = errors.New("cannot divide by zero")
		return res, rem, err // zero value
	}
	res, rem = num/den, num%den
	return res, rem, err
}

func divAndRemainder2(num int, den int) (res int, rem int, err error) { // parentheses, even if it is one
	// A problem with named return variables is that they may not returned
	res, rem = 20, 30
	if den == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / den, num % den, nil
	// The Go compiler inserts code that assigns whatever is returned to the return parameters
	// They provide a better documentation, but they are of limited value and shadowing is possible.
}

//Blank returns (never use them)
// only write return and the compiler return the last values assigned to the named return values
func divAndRemainder3(num int, den int) (res int, rem int, err error) {
	if den == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	res, rem = num/den, num%den
	return
}

// When there is invalid input, we return inmidiately, zero values are return but they may not make sense
// They are a bad practice

// Signature of the function
// The type of a function is  built out of the keyword func and the types of the parameters and returns
// Let's create a calculator

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

// Funcion type declarations
// A function can be a type and be defined
type opFuncType func(int, int) int //comment

// for example rewriting the opMap:
var opMap = map[string]opFuncType{} //comment

func main() {
	x, y, z := divAndRemainder(5, 2)
	fmt.Println(x, y, z)
	x2, y2, z2 := divAndRemainder2(5, 2)
	fmt.Println("func 2 ", x2, y2, z2)
	x3, y3, z3 := divAndRemainder3(5, 2)
	fmt.Println("func 3 ", x3, y3, z3)

	// Calling the calculator
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("INVALID EXPRESISON")
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)

	}

}
