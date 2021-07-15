package main

import "fmt"

func main() {

	// Booleans
	var flag bool // default is false
	var isAwesome = true
	// Only with fmt.Printf we can print the variable with %v
	fmt.Printf("flag is %v\n", flag)
	fmt.Printf("is it Awesome? %v\n", isAwesome)

	// Integers
	var int1 int = 2
	var int2 int32 = 42
	var int3 int8 = -128
	fmt.Printf("Value is: %v and type is: %T\n", int1, int1)
	fmt.Printf("Value is: %v and type is: %T\n", int2, int2)
	fmt.Printf("Value is: %v and type is: %T\n", int3, int3)
	// fmt.Printf("int + int64", int1+int4)
	// uints are only positive numbers
	var uint1 uint8 = 255
	fmt.Printf("Value is: %v and type is: %T\n", uint1, uint1)

	// floats are not exact in their memory representation
	var fl1 float64 = 3.124411
	// conversions have to be explicit even between ints and floats.
	// float64 only with float64 and the same with other types, even int and int64.
	fmt.Printf("The stored value is: %v\n", fl1)
	fmt.Printf("int64 converted plus float64: %v\n", fl1+float64(int1)) // Conversion types

}
