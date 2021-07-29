package main

import(
	"fmt"
	"io"
)

// The real star of Go design are interfaces, the only abstract type in Go.
type Stringer interface { // declaration, appears adter the name of the interface type
	// Stringer interface of the fmt package
	// It list the methods that must be implemented by a concrete type
	// Methods defined by an interface are called the method set of interface
	String() string
} // They are named with "er"

// The concrete type (it implements all methods from the interface) can be assinged to a variable
// or field declared to be of the type of the interface

// types in Go enable both type-safety and decoupling

// Program to an interface, not an implementation, it makes dependance in behaviour, so
// swap implementations can be done

type LogicProvider struct{} // no indication it implements an interface, code can
// evolve, maybe another LogicProvider will be needed, but the Client may not change

func (lp LogicProvider) Process(data string) string {
	//business logic
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic // Because itimplements the Logic interface, when creating a Client,
	// it need a type of LogicProvider
}

func (c Client) Program() {
	//get data from somewhere
}

// Interfaces specify what callers need, client code defines the interface 
// to specify what functionality it requieres 

// interfaces can be shared, using standard interfaces encourages the decocrator pattern
// In Go is common to write factory functions that takes in an instance of an interface and
// return another type that implements the same interface 

func process(r io.Reader) error
// process data from a file:
r, err := os.Open(filename)
if err != nil {
	return err
}
defer r.CLose()
return process(r)
return nil

// r is an os.File instance, returned by os.Open. It meets the io.Reader interface and can be
// used in any code that reads in data. For example, let's read a compressed file:

r, err := os.Open(filename)
if err != nil {
	return err
}
defer r.Close()
gz, err = gzip.NewReader(r)
if err != nil {
	return err
}
defer gz.Close()
return process(gz)
// The same code reading from an uncompressed, now it reading from a compressed file.

// It's fine for a type that meets an interface to specify addional methods that aren't part of the interface.


func main() {

	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
