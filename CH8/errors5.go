package main

import (
	"fmt"
	"os"
)

// Wrapping errors with defer
// Wrapping multiple errors with the same message:

func DoSomething(val1 int, val2 string) (string, error) {
	val3, err := doThing1(val1)
	if err != nil {
		return "", fmt.Errorf("in DoSomehing: %w", err)
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", fmt.Errorf("in DoSomehing: %w", err)
	}
	result, err := doThing3(val3, val4)
	if err != nil {
		return "", fmt.Errorf("in DoSomehing: %w", err)
	}
	return result, nil
}

// Simplifing with defer

func DoSomething(val1 int, val2 string) (_ string, err error) {
	// name return values so that we can refer to err in te deferred func
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomething_ %w", err)
		}
	}()

	val3, err := doThing1(val1)
	if err != nil {
		return "", err
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", err
	}
	return doThing3(val3, val4)
}

// Panic and defer
// Go generates a panic whenever there is a situation whertr Go runtime is
// unable to figure out what should happen next.
// The order is as follows: panic -> defer -> ... -> main

func doPanic(msg string) {
	panic(msg)
}

// recover gets the panic message and continue the execution of the script
func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

// There is a specific pattern for using recover. In te ex above, we registered a
// function with defer to handle a potential panic. You must call recover withing
// a defer, because once a panic happends, only deferred functions run

// recover  does not make clear what could fail, only prints an error and

// A common error is wanting a stack trace when something goes wrong, Go does not
// support it, it is best to wrap error with stack traces

func main() {
	doPanic(os.Args[0])

	// panic: /tmpfs/play
	// goroutine 1 [running]:
	// main.doPanic(...)
	// 	/tmp/sandbox567884271/prog.go:6
	// main.main()
	// 	/tmp/sandbox567884271/prog.go:10 +0x5f

	for _, val := range []int{1, 2, 0, 9} {
		div60(val)
	}
	// 60
	// 30
	// runtime error: integer divide by zero
	// 10
}
