package main

import (
	"fmt"
	"io"
)

// Good practice to use type assertions
func doThings(i interface{}) {
	//type switch, specifying a variable of type interface, followed by .(type)
	switch j := i.(type) {
	// Type switch purpose is to derive a new variable from an existing one, shadowing
	// is idiomatic, but it is not used in this example, should have been: (i := i.(type))
	// Type switch provide the ability to differentiate between multiple implementations of an interface
case nil:
		// i is nil, type of j is interface
	case int:
		// j's type is int
	case MyInt:
		// j's type is MyInt
	case io.Reader:
		// j's type is io.Reader
	case string:
		// j's type is string
	case bool, rune:
		// i is either a bool or rune, so j is of type interface{}
	default: // it is important the default, this handles unknown implementations in devolopment stage
		// no idea what i is, so j is of type interface{}
	}
}

type MyInt int

// Use type assertions and type witches sparingly
// One common use of type assertion is to see if the concrete type behind an interface,
// implements another interface. This let specify optional interfaces

// this function is the actual implementation of Copy and CopyBuffer
// if buf is nil, one is allocated
func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy
	// Avoids an allocation and a copy
	if wt, ok := src.(WriterTo); ok { // If it implements WriterTo, a lotofwork can be skipped
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has ReadFrom method, use it to do the copy
	if rt, ok := dst.(ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
	// func continues
}

// Another place optional interfaces are used is when evolving an API, with "context", but it is from go >= 1.7

// If there is an optional interface implemented by one of the wrapped implementations, you
// cannot detect it with a type assertion or type switch

func walkTree(t *treeNode) (int, error) {
	switch val := t.val.(type) {
	case nil:
		return int(val), nil
	case number: // we know t.val is of type number, return int value
		return int/(val), nil
	case operator:
		// We know that t.val is of type operator, so find the values of the left and right children,
		// then call the process() method on operator to return the result of processing their values
		left, err := walkTree(t.lchild)
		if err != nil {
			reurn 0, err 
		}
		right, err := walkTree(t.rchild)
		if err != nil {
			reurn 0, err 
		}
		return val.process(left, right), nil
	default:
		//If a new treeVal type is defined, but walkTrre wasn't updated to process it, this detects it:
		return 0, error.New("Unknown node type")
	}
}

func main() {

	// Type assertions and type switches
	// They are used for reading back the value of an interface by checking a
	// concrete type or if it implements another interface.

	// Type assertion from MyInt
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt) // .(type) to do a type assertion, it makes it MyInt
	fmt.Println(i2 + 1)
	// i3 := i.(string)
	// fmt.Println(i3 + 1) // the interface is now string, error
	// i4 := i.(int)
	// fmt.Println(i4 + 1) // even thought the underling type is int, causes an error

	// To avoid crashing
	i5, ok := i.(int) //
	if !ok {
		fmt.Errorf("unexpected type for %v", i)
	}
	fmt.Println(i5 + 1)

	// Type conversions can be applied to concrete types and interfaces, checked at compilation time
	// Type assertions only for interfaces and are checked at runtime (so they can fail)

}
