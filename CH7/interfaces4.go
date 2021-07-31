package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Function types are a bridge to interfaces
// Go allows methods on any user-defined type, including user-defined function types,
// this allow functions to implement interfaces and the most common usage is for HTTP
// handlers, these processes an HTTP server request. It's defined by an interface:
type Handler interface {
	ServerHTTP(http.ResponseWriter, *http.Request)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

// By using type conversion to http.HandlerFUnc, any functoin that has the same signature
// can be used as an http.Handler

func (f HandlerFunc) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

// This lets you implement HTTP handlers using functions, methods, or closure using the exact
// code path as the one used for other types that meet the http.Handler interface

// If your single function is likely to depend on many other functions or other state that's not
// specified in its input parameters, use an interface parameter and define a function type to
// bridge a function to the interface

// Implicit interfaces make dependency injection easier (decoupling), your code should
// should explicitly specify the functionality it needs to perform its task

// Simple web app
// Logger function
func LogOutput(message string) {
	fmt.Println(message)
}

// Data store
type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
	// This method meets the interface DataStore
}

// Factory function to create an instance of SimpleDataStore
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Bob",
			"3": "Pat",
		},
	}
}

// Businesslogic
// It requires data to work with (a data store), logging when invoked (depends on a logger)
// but we don't want to make it dependable of LogOutput or SimpleDataStore

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}
type Logger interface {
	Log(message string)
}

// To make our LogOutput function meet this interface, we define a function type with a method on it

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
	// This method meets the interface Logger
}

// Bussines logic implementation
type SimpleLogic struct {
	// interface fields, no mention of concrete types, so, no dependency on them. No problrm
	// if we later swap in new implementationa from an entirely different provider
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

// When we want a SimpleLogic instance, we call a factory function
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	// passing interfaces and returning a struct
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// To our API, we only need a single endpoint, /hello, which says hello to the
// person whose user ID is supplied. Our controller needs business logic:
type Logic interface {
	SayHello(userID string) (string, error)
	// THis method is avaible on our SimpleLogic struct, but the concrete type
	// is not aware of the interface
}

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In Sayhello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

// Just as we have factory functions for our other types, let's write one for the Controller
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func main() {
	// Wiring up all of our components for the web app
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)

}
