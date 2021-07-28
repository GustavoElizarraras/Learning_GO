package main

import (
	"fmt"
)

// Embedding for composition
type Employee struct {
	Name string
	ID   int
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%v)", e.Name, e.ID)
}

type Manager struct {
	Employee // Contains field of type Employee, without name, so it is a embedded field.
	// This means any fields or methods declared on an embedded field are promoted to the containing struct
	// and can be invoked directly on it.
	Reports []Employee
	// You can embed any type within a struct, except another struct. This promotes the methods on the
	// embedded type on the containing struct
}

// func (m Manager) FindNewEmployees() []Employee {
// 	//LOGIC
// }

func main() {
	// in other languages, enumarations let specify that a type can only have a limited set of values
	// In Go, we have iota lets assign an increasing value to a set of constants
	// When using iota, the best practice is to first define a type based on int that
	// will represent all of the valid values.
	type MailCategory int
	// Now use a const block to define a set values for your type:
	const (
		Uncategorized MailCategory = iota // specified type and value set to iota
		// Neither value or type:
		Personal
		Spam
		Social
	)
	// When the compiler see this, it repeats the type and increments the value of iota

	// Use iotaonly for internal purposes, where constants are referred to by name rather than value,
	// this way we can enjoy iota by inserting new constants at any moment in time/location in the list
	// without the risk of breaking everything.
	// iota based enumerations only make sense when differenting between a set of values and the value is
	// not important. If the value matter is better tzo be explicit.

	type BitField int
	const (
		F1 BitField = 1 << iota // this is clever, but a behaviour like this has to be documented
		F2                      // 2
		F3                      // 4
		F4                      // 8
	)

	// iota starts numering from 0. If there isn't a sensical default value, assing _ to the first iota to indicate it is invalid.

	// Embedded
	m := Manager{
		Employee: Employee{
			Name: "Gus",
			ID:   276,
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	fmt.Println(m.Description())

	// We cannot assign a variable of type Manager to a var of type Employee, we have to be explicit about
	// accesing the Employee field in Manager
	// var eFail Employee = m // error
	var eOK Employee = m.Employee
	fmt.Println(eOK)
}
