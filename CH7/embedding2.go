package main

import (
	"fmt"
)

// If the containing struct has fields or methods with the same name as  an embedded field,
// we need to use the embedded field's type to refer to the obscured fields or methods
type Inner struct {
	x int
}
type Outer struct {
	Inner
	x int
	s string
}

// No dynamic dispatch
// It is something alike with methods, if we havea method on an embedded field that calls another
// another method on the embeded, and the containig struct has a method with the same name, the embedded
// method will not invoke the method on the containing struct

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}
func (i Inner) Double() string {
	return i.IntPrinter(i.x * 2)
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func (o Outer) Double() string {
	return o.IntPrinter(o.x * 2)
}

func main() {
	o := Outer{
		Inner: Inner{
			x: 10,
		},
		x: 20,
		s: "Hello",
	}
	fmt.Println(o.x)       // Outer
	fmt.Println(o.Inner.x) // Inner, specifiing the embedded field type

	// Methods
	fmt.Println(o.Double())       // If the method Double that receives an Outer didn't exist, it would call the Inner
	fmt.Println(o.Inner.Double()) // If we want an the Double method that receives the an Inner type, we have to be explicit

}
