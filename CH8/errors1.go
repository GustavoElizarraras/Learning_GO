package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
)

// Go handles errors by returning a value of type error as the last
// return value for a function it is a strong convention. When no error
// is present, nil is return for the error parameter

func calcRamainderAndMod(num, den int) (int, int, error) {
	if den == 0 {
		return 0, 0, errors.New("denominator is 0") // No capitalized, no \n or punc
	}
	return num / den, num % den, nil
}

// error is a built-in interface that defines a single method:
type error interface {
	// Everything that implements this interface is considered an error
	// We return nil because is the zero value for an interface type
	Error() string
}

// Go uses a return error instead of thrown exceptions, because this leads to clear and
// more stable code and because the Go compiler requieres all variables to be used, so if
// an error var is meant to be ignored, the _ is used

// Use strings for simple errors, with errors.New(), fmt.Errorf()

// Sentinels error are used to indicate that a procces can't start or continue because a
// problem with the current state. Their name start with Err

// Be sure to need a sentinel error, when one is defined, it becames part of your public API,
// is far better to rehuse one in the standard library or an error type with information. But
// if a specific state has been reached where no futher processing is possibñe, use a sentinel

// Do not use constants for sentinel errors:

// package consterr

// type Sentinel string

// func(s Sentinelm) Error() string {
// 	return string(s)
// }

// package mypkg

// const(
// 	casting a string literal to a type that implements the error interface
// 	ErrFoo = consterr.Sentinel("foo error")
// 	ErrBar = consterr.Sentinel("bar error")
// 	If you used the same type to create constant errors across packages, two errors
// 	would be equal if the strings are equal. It is better to use errors.New()
// )

func main() {

	// Go does not ahve special constructs to detect if an error was returned, check with an if
	n, d := 20, 3
	remain, mod, err := calcRamainderAndMod(n, d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(remain, mod, err)

	// ex of sentinel error
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	no, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		// test if the error returned is a sentinel error (it is said in the documentation)
		fmt.Println("Told you so", no)
	}
}
