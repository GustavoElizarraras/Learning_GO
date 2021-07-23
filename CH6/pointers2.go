package main

import "fmt"

// Pointers are a last resort, they could make harder to understand dataflow and create extra work for the garbage collector
// It is more complicated populating a struct by passing a pointer to it into a function
type Foo struct {
	Field1 string
	Field2 int
}

func MakeFooBad(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 2
	return nil
}

// It is easier to just have the function instantiate and return the struct
func MakeFoo() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 23,
	}
	return f, nil
}

// When returning values from a function you should favor value types, only use pointer types
// if there is state within the data type that needs to be modified. It is commonly used in buffers and concurrency.

// If the struct or data structure is large enough (10MB), it is faster using pointers than a value type

// Resist the temptation  to use a pointer field to indicate no value, is better to have a value type paired with a boolean.

// Passing a map to a function, you pass a copy of the pointer that reference the map, bad idea to use them in APIs because it is not explicit
// what keys are in the map, so, it is not self-documenting.

func appendSlice(x_copy []int) []int {
	x_copy = append(x_copy, 4)
	return x_copy
}

func addValue(x_val []int) {
	x_val[4] = 1
}

func main() {

	bad := Foo{}
	fmt.Println(MakeFooBad(&bad))
	fmt.Println(MakeFoo())

	x := make([]int, 3, 10)
	fmt.Println(x)
	y := appendSlice(x)
	fmt.Println(x)
	fmt.Println(y)
	// addValue(x) // error, different length sizes, the original can't see it
	// fmt.Println(x)

	// Go encourages to use pointers sparingly, to reduce the workload of the garbage collector
	// by making sure that as much data is stored on the stack. For example, slices of structs or
	// primitive types have their data lined up sequentialy in memory for rapid access
}
