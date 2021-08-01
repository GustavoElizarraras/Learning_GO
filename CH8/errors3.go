package main

import(
	"fmt"
	"errors"
	"os"
)

// Wrapping Errors 
// Add additional context to an error, it can be the name of the function
// that received the error or the operation it was trying to perform. This
// is called wrapping the error, and if there is a series of it, it is an error chain

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		// Wrapping the error with Errorf, it has the special verb %w, use it
		// to create an error with a formatted string + original error string 
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil 
}

// If you want to wrap an error with your custom error type, your error needs
// to implement the method Unwrap

type Status int 

type StatusErr struct{
	Status Status
	Message string
	err error 
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.err
}

// Now we can use StatusErr to wrap underlying errors

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err := nil {
		return nil, StatusErr{
			Status: InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid)
			Err: err,
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status: NotFound,
			Message: fmt.Sprintf("file %s not found", file),
			Err: err,
		}
	}
	return data, nil
}

// A libray may return an error meaning the process cant continue, but that error contains
// implementation details that are not needed for other parts of the code, in this case, 
// a new brand error may be created 

// If you want to create a new error that contains the message of another, without wrapping it,
// use fmt.Errorf(with %v instead of %w)

err := internalFUnction()
if err != nil {
	return fmt.Errorf("internal failure %v", err)
}

func main () {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		// errors.Unwrap() returns the wrapped error, if there is one, if not, nil
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
}