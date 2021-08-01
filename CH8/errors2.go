package main

import (
	"fmt"
)

// Errors are values
// error is an interface, so you can define your own errors that include additional
// information for logging or error handling

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string { // holding StatusErr value
	return se.Message
}

// using StatusErr to provide more details about what it went wrong

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	// even if custom error type, always use error as the return type for the
	// error result. This allows to return different types of errors from the
	// function and allow callers of the func to not depend on a specific error type
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),QQQQQA
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

// If using a custom error, make sure you don't return an uninitialized instance
func GenerateError(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr 
	// for an instance to be considered nil, the type and underlying value must be nil
	// the underlying type part of the interface is not nil
}

// Way 1 to correct this: return nil explicitly when the function completes succesfully

func GenerateErrorGood1(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	return nil
}

// Way 2: make sure the local variable holding an error is of type error

func GenerateErrorGood2(flag bool) error {
	var genErr error
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr 
}

// When using custom errors, never define a varible to be of the type of your custom error
// either explicitly return nil when no error occurs or define the variable to
// to be of type error

func main() {
	err := GenerateError(true)
	fmt.Println(err != nil) // true
	err = GenerateError(false)
	fmt.Println(err != nil) // true
}
