package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

// Is and As
// If a sentinel error is wrapped, you cannot use == to check for it, nor type
// asserition and type switch to match a wrapped custom error. Is and As solve this

// Use is to check if an error wraps a specific sentinel error
func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

// If errors.Is() does not work for a custom error, implement that method

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return reflect.DeepEqual(me, me2) // This can compare anything, even slices
	}
	return false
}

// Making a custon Is() method allow comparisons against errors that aren't identical
// instances, maybe pattern match errors, specifying a filter instance that matches
// errors that have some of the same fields

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

// If we want two ResourceErr instances to match when either field is set, let's do a custom Is()
func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code
		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
}

// Now we can find, for example, all errors regarding the database, no matter the code
// if errors.Is(err, ResourceErr{Resource: "Database"}) {
// 	fmt.Prinln("The database is broken:", err)
// 	// process the codes
// }

// error.As() function allows to check if a return error (or anyerror it wraps) matches a specific type
ferr := AFunctionThatReturnsAnError()
var myErr MyErr
if errors.As(ferr, &myErr) {
	// First is the error to be examined, second is a pointer to a variable of the search type returns
	// true if an error in the chain was found, and if it is found, is assingned to the pointer
	fmt.Println(myErr.Code)
}
// Or maybe the second parameter can be an interface, to find an error that meets it
var coder interface {
	Code() int 
}
if errors.As(ferr, &coder) {
	// Any interface type is acceptable
	fmt.Println(coder.Code())
}

// It is possible to override the As method creating a custom, but it would be very unnusual,
// maybe to match an error of one type but return another

// Use Is() to look for a specific instance or value, use As to look for a specific type

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// by default errors.Is uses == to compara each wrapped
			// method with the specified one
			fmt.Println("Thatfile does not exist")
		}

	}
}
