package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Embedding and Interfaces
// We can embed an interface in an interface, or embed an interface in a struct
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadCLoser interface { // built on of the Reader and Closer interfaces
	Reader
	Closer
}

// Accept interfaces, return structs
// // The business logic invoked of functions should be invoked via interfaces
// // but the output should be a concrete type

// // When creating an API, returning interfaces result in coupling, the code is dependant
// // on the module that contains those interfaces and dependencies. To docouple, another interface
// // is needed and do types conversions

// // Versioning, if a concrete type is returned, new methods and fields can be added without refactoring,
// // if returning an interface, adding a new method means updating all implementations of the inteface

// // Rather than writing a single factory function that returns different instances behind an interface
// // based on input parameters, try to writeseparate factory functions for each concrete type

// // Errors are the exception, Go functions and methods declare a return parameter of the error interface type,
// // we need to use an interface to handle all posible options.

// // When invoking a function with parameters of interface types, a heap allocation occurs for each
// // of the interfce parameters, returning a struct avoids this; so, we need to figure out the trade off
// // between better abstraction and better performance. In general, write mantainable and readable code.

func main() {

	// interfaces and nil
	// // nil is the value type for pointer types
	// // For an interface to be nil, it's type and value must be nil
	// // In Go runtime, interfaces are implemented as a pair of pointers, one for type and the other for value
	// // What nil means for interfaces is weather or not methods can be invoked on it. If an
	// // interface is nil, invoking any methods on it triggers a panic. Te nil value for interfaces
	// // should be implemented/handled in methods.
	var s *string
	fmt.Println(s == nil)
	var i interface{}
	fmt.Println(i == nil)
	i = s
	fmt.Println(i == nil) // false

	// Empty interface says nothing
	// // interfaces can store a value of any type
	var j interface{} // can be used as placeholder for data of uncertain schema from an external source
	j = 20
	fmt.Println("Hi, i'm an interface and my value is:", j)
	j = "hello"
	fmt.Println("Hi, i'm an interface and my value is:", j)
	j = struct {
		a int
		b string
	}{1, "b"}
	fmt.Println("Hi, i'm an interface and my value is:", j)

	// json example
	// one set of braces are for the nil interface{}, the other for an instance of the map
	// // data := map[string]interface{}{} 
	// // contents, err := ioutil.ReadFile("file.json")
	// // if err != nil {
	// // 	return err
	// // }
	// // defer contents.Close()
	// // json.Unmarshal(contents, &data) // contents in data map

	// Another implementation is for user-created data structures, work with multiples types and to hold any value
	// // type LinkedList struct {
	// // 	Value interface{}
	// // 	Next *LinkedList
	// // }

	// // func (ll *LinkedList) Insert(posi int, val interface{}) *LinkedList {
	// // 	if ll == nil || pos == 0 {
	// // 		return &LInkedList{
	// // 			Value: val,
	// // 			Next: ll,
	// // 		}
	// // 	}
	// // 	ll.Next = ll.Next.Insert(posi-1, val)
	// // 	return ll 
	}
	// In general, avoid using interface{}, Go is strongly typed and attempts to work around this are unidiomatic
}
