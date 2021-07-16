package main

import "fmt"

func main() {

	// Structs are used for related data of different type
	// Also, they are offer OOP capabilities, Go does not have classes, so, nor inheritance

	type person struct { // Declaring a struct
		// Fields
		name string
		age  int
		food string
	}
	// They can be outside or inside a function, if it is inside, it can only be used within thaf func

	var fred person // No value assigned, zero value struct (all values are zero)
	fmt.Printf("fred is: %v\n", fred)

	// A struct literal can be assigned to a varaible
	bob := person{} // Empty struct literal, also zero values
	fmt.Printf("bob is: %v\n", bob)

	// Structs with values
	moni := person{
		// ordered as the struct is created
		"Moni",
		48,
		"avocado",
	}

	saul := person{
		// specify the fields in any order, better this way
		age:  45,
		food: "chicharrón",
		name: "Saúl",
	}

	fmt.Printf("Moni: %v\nSaúl: %v\n", moni, saul)

	// If initializing a struct without the name fields,
	// and a future version implements additional fields, it will no longer compile

	fmt.Printf("Moni's fav food is: %v\n", moni.food) // accesing fields using dot notation

	// Anonymous structs, a variable implements a struct type
	// They are a single instance, used in json and table-driven test
	var taco struct {
		meat string
		many int
	}
	taco.meat = "pork"
	taco.many = 5
	fmt.Printf("Order: %v\n", taco)
	taco.many = 6 // Internals are mutable
	fmt.Printf("Order: %v\n", taco)

	// Pair of brackets to assign the fields
	cellphone := struct {
		brand string
		year  int
	}{
		brand: "Apple",
		year:  2021,
	}
	fmt.Println(cellphone)

	// Comparing structs
	// Depend if the fields are comparable, slices and maps are not
	// Gp does not allow comparisons between variables that represent structs of different types
	// Type conversions between structs ONLY IF the fields have the same names, order and types
	// Comparison is allowed between named and anonymous structs if fields are identical

	type firstPerson struct {
		name string
		age  int
	}
	one := firstPerson{
		"One",
		1,
	}
	type secondPerson struct {
		name string
		age  int
	}
	two := secondPerson{
		"Two",
		2,
	}
	// one==two is not allowed
	fmt.Printf("type fp: %T and type sp: %T\n", one, two)
	fmt.Printf("type fp: %T\n", secondPerson(one)) // one is now a secondPerson type

	fmt.Printf("Is one the same as two?: %v\n", secondPerson(one) == two) // same type, different field values; so false

}
