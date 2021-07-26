package main

import (
	"fmt"
)

// Code methods for nil instances
// Go will try to invoke the method, if it has a value receiver, a panic will be raised,
// if it has a pointer receiver and it is written to handle a nil instance, it will work
// Binary tree implementation

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool { // doesn't modify the *IntTree
	switch {
	case it == nil: // Supports a nil receiver
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

// You can't write a pointer receiver method that handles nil and makes the
// original non-nil because it's a copy of the pointer that's passed into the method.

// Methods are functions too
// functions can be replaced anytime there's a variable or parameter of a function type
type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

// If the function depends depends on other data, when startup values change while the programm
// is running, those changes should be stored in a struct, and the logic implementation should be a
// method. Otherwise, if the logic depends on input parameters, a function is good to go

func main() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))
	fmt.Println(it)

	// Invoking methods
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))
	// Assign method to variable and pass a type func(int) int. Method value
	f1 := myAdder.AddTo // Similar to closure, it access the values in the fileds
	// of the instance from which it was created
	fmt.Println(f1(10))
	// Create a function from the type itself. Method expression
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15)) // The first parameter is the receiver for the method

	// Inheritance does not exist in Go, we can declare user-definied types
	type Score int
	type HighScore Score
	// They only have the underling type, no hierarchy, no shared methodss, always use a type conversion
	// The built in operators can be used in user defined types built from them.
	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	// hs = s // error
	// s = i // error
	fmt.Println("score:", s, "high score:", hs)
	s = Score(i)
	hs = HighScore(s)
	fmt.Println("score:", s, "high score:", hs)
}
