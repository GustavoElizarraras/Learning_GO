package main

import (
	"fmt"
	"time"
)

// Methods
// Access directly a field instead of writing getters and setters, exept when meeting an interfece or
// when updating multiple fields as a single operation or when the update is not a straighforward assigment
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string { // the receiver is between the keyword func and the name of the method
	// The receiver name is a short abbreviation of the type's name, usually its first letter. It is
	// non-idiomatic using this or self
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// Not reusing names is a good practice in Go, it make clear what the code does
// Keep together the implementation

// Pointer Receiver and Value Receivers
// // If a method modifies the receiver, a pointer receiver must be used
// // If a mehod need to handle nil instances, a pointer receiver must be used
// // If a method does not modify the receiver, a value receiver can be used
// // If a type uses pointer receiver methods, it is a better practice to use them
// // in all methods, even the ones that dont modify the receiver.
type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}
func (c *Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}
func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("Updating right", c.String())
}
func doUpdateWrong(c Counter) { // In the wrong one, we are passing a value instance
	c.Increment()
	fmt.Println("Updating wrong", c.String())
}

func main() {

	// In addition to struct literals, types can be composed using primitive types or compound type literals
	type Score int
	type Converter func(string) Score
	type TeamScores map[string]Score
	// Abstract types: specifies wha a type should do, but not how to
	// Concrete types: specifies what and how, specific ways to store data and provides an implementation of
	// any methods declared on the type

	// Summon the Person String method
	p1 := Person{
		FirstName: "Gustavo",
		LastName:  "Elizarraras",
		Age:       52,
	}
	fmt.Println(p1.String())
	// Trying out pointer receiver method
	var c1 Counter
	fmt.Println(c1.String())
	c1.Increment()
	// We are passing a value type and calling a pointer receiver method on a copy
	fmt.Println(c1.String())
	// c1 is a value type, when using a pointer receiver with a local variable that's a value type, Go converts it
	// to a pointer type (c) -> (&c)

	// Understanding better with another example, dopUpdateRight and doUpdateWrong
	var c2 Counter
	fmt.Println(c2.String())
	doUpdateRight(&c2)
	fmt.Println(c2.String())
	doUpdateWrong(c2)
	fmt.Println(c2.String())

}
