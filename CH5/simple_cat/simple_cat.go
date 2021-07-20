package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 { // checking specified name of file
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err) // If missing argument, print a message and exit
	}
	// defer is used for cleaning temporary resources
	defer f.Close() // need to close file after we use it,
	// defer delays the function invocation until the sorrounding function exits
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data) // returns the number of bytes that were read and an error
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

	// We can defer multiple closures in a Go function. They run in last-in-first-out order,
	// so the last defer registered runs first.

	// The code within defer runs after the return statement

	// Any variable passed into a deferred closure are not evalued until the closure runs.

	// A way for a deferred function to examine or modify the return values of its surrounding funtion
	// is using named return values, these take actions based on an error.

}
