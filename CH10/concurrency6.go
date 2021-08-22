package main

import (
	"sync"
)

// Using WaitGroups
// Sometimes one goroutine need to wait for multiple goroutines to complete
// their work, if it is only one, the done channel pattern is enough; if it
// depends on multiple, we need to use a WaitGroup

func ex1() {
	// a sync.WaitGroup does not need to be initialized, just declared, as
	// its zero value is useful
	var wg sync.WaitGroup // We does not pass it explicitly, we assure is
	// the same instance (a pointer should work too), we could use a closure
	// to ensure it is the same instance. The usual pattern is to launch a
	// goroutine with a closure that wraps business logic, this handle
	// concurrency and provides the algorithm.

	wg.Add(3) // increments the counter of goroutines to wait for
	go func() {
		// Done() decrements de counter, it is called inside the goroutine
		defer wg.Done()
		doThing()
	}()
	go func() {
		defer wg.Done()
		doThing2()
	}()
	go func() {
		defer wg.Done()
		doThing3()
	}()
	wg.Wait() // pauses its goroutine until the counter hits zero

}

// If multiple goroutines write to the same channel, it only need to be closed once
// ex2, we launch a monitoring goroutine that waits until all of the processing
// goroutines exit
func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []int
	for v := range out {
		// This loop exits when out is closed and the buffer is empty
		result = append(result, v)
	}
	return result
}

// WaitGroups are handy, but, use them only when you have something to clean up
// (like closing a channel they all write to) after all of the worker goroutines exit

// Running code exactly once
// init should be reserved for initialization of effectively immutable package-level state,
// but sometimes, lazy load is requiered, call some initialization code exactly once after
// program launch time. This is usually because the initialization is relatively slow and
// may not even be needed every time your program runs, Once enable this functionality

type SlowComplicatedParser interface {
	Parse(string) string
}

// package level variables
var parser SlowComplicatedParser

// we don't have to configure this instance, making the zero-value useful:
var once sync.Once // make sure not to copy, don't declare it inside a func

func Parse(dataToParse string) string {
	// if Parse is called more than once, once.Do() will not execute again
	once.Do(func() {
		// parser is only initialized once, inside this closure
		parser = initParser()
	})
	return parser.Parse(dataToParse)

}

func initParser() SlowComplicatedParser {
	// do all sorts of setup and loading here
}

func main() {

}
