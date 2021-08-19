package main

import (
	"fmt"
)

// Concurrency practices and patterns

// A good API implementation, allows to change how to code works without
// changing how is invoked. So, you should never expose channels or
// mutexes in your API's types, functions and methods. If a channel
// is exposed, the user now has to manage it, take care of deadlocks, etc.
// You can have them as function parameters or struct fields, but never
// export them. The exception is if the API is a library with a
// concurrency helper function, like time.After

// Goroutines, for Loops, and Varying variables
// Sometimes the closure that launches a goroutine has no parameters,
// instead it captures values from the enviroment. This does not work
// with index or value of a for loop

func ex1() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		// The index and value variables in a for loop are rehused in each
		// iteration, in this case, the last value assigned was 10
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
	// prints: 20 20 20 20 20
}

for _, v := range a {
	// A way to fixed this is to shadow the variable inside the loop
	v := v
	go func(){
		ch <- v * 2
	}()
}

for _, v := range a {
	// Another is to pass it to the goroutine
	go func(val int){
		ch <- val * 2
	}(v)
}

// Always clean up your goroutines
// Go runtime can't detect that a goroutine will never be used again, so
// if it doesn't exit, the scheduler will still periodically give it time
// to do nothing, so the programm will be slowed, this is a goroutine leak

func countTo(max int) <-chan int {
	ch := make(chan int)
	// goroutine as a generator: (bad example, simple task)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// The done channel pattern
// Is a signal to a goroutine tha it's time to stop processing, it uses a
// channel to do it

func searchData(s string, searches []func(string) []string) []string {
	// Channel cantains an empty struct because the value is indiferent,
	// we never write on it, only close it
	done := make(chan struct{}) 
	// We read the first value written to result, and then we close done.
	// This signals to the goroutines that they should exit preventing them
	// from leaking
	result := make(chan []string)
	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			// wait for a write in the result channel or a read on the done ch
			select {
			case result <- searcher(s):
			case <- done:
			// this case, reads from done and will stay paused until it is closed
			}
		}(searcher)
	}
	r := <-result 
	close(done)
	return r
}


func main() {

	for i := range countTo(5) {
		// In a common case, where you use all of the values, the goroutine
		// exits, if we exit the loop early, the goroutine blocks forever,
		// waiting for a value to be read from the channel
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}
