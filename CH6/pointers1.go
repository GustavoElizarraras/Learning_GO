package main

import (
	"fmt"
)

// helper function that takes in a primituve and return a pointer to that type
func stringp(s string) *string { return &s }

// tip: use a helper function to turn a constant value into a pointer

// functions to observe pointers behaviour
func failedUpdate(px *int) {
	x2 := 20
	px = &x2 // change the copy, not the original
}

func update(px *int) {
	x2 := 20
	*px = x2 // dereferencing points to the original and the copy
}

func main() {
	// Every variable is stored in one or more contiguous memory locations, called addresses
	// All pointers are of the same size
	// The zero value for a pointer is nil

	// The & is the address operator, precedes a value type and returns the address of the memory location
	x := "hello"
	pointerToX := &x
	// The * is the indirection operator, precedes a variable of pointer type and returns the pointed-to value.
	// This is called dereferencing
	fmt.Println("x is: ", x)
	fmt.Println("x's pointer is: ", pointerToX)             // address memory
	fmt.Println("the value of pointeToX is: ", *pointerToX) // what is located in the address
	// Before dereferencing a pointer, check it is non-null
	var nilPointer *int
	fmt.Println(nilPointer == nil)
	// fmt.Println(*nilPointer) //panics

	// A pointer type is a type that represents a pointer. It is written with a * before a type name
	y := 10
	var pointerToY *int
	pointerToY = &y
	fmt.Println("y is: ", *pointerToY)

	// new() function creates a pointer variable with zero value. It is rarely used
	var newVar = new(int)
	fmt.Println("Is the variable created with new() nil?", newVar == nil)

	// For structs use an & before a struct literal to create a pointer instance
	type Foo struct {
		i int
		s string
	}
	stru := &Foo{}
	stru.s = "a"
	fmt.Println(stru)

	// When you need a pointer to a primitive type, declare a variable and then point to it
	var a string
	b := &a
	fmt.Println(b)

	// If a struct has a field of a pointer to a primitive type, a literal can't be assigned to the field
	type person struct {
		FirstName  string
		MiddleName *string // The same workaround of the previous lines, holding variable or helper function
		LastName   string
	}

	p := person{
		FirstName: "Ale",
		// MiddleName: "Eli", //with & or without, it won't compile
		MiddleName: stringp("No middle name"),
		LastName:   "Lla",
	}
	fmt.Println(p, *p.MiddleName)

	// Pointers indicate mutable parameters, we can choose between value and pointer parameter types
	// With pointers, a parameter is mutable
	// For non-pointers types like primitives, structs and arrays, the inmutability of the original
	// data is guaranteed because go is a call for value language. But if a pointer is passed to the function,
	// it gets a copy of the pointer, so, the original data can be modified
	d := 10
	failedUpdate(&d)
	fmt.Println(d)
	update(&d)
	fmt.Println(d)
}
