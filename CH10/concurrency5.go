package main

import (
	"errors"
	"time"
)

// Turning off a case in a select
// selectis used to combine data from multiple concurrent sources, but you
// need to properly handle closed channel. If a case is is reading a closed
// channel, it will return the zero value, you need to validate the value and 
// skip the case, if reads are spaced out, the programm will waste time reading junk
// We can disable a case in a select with a nil channel, when a channel has been
// closed, set the channel variable to nil, the associated case will no longer run
// because the reading from a nil channel never returns a value  

for {
	select {
	// process the v that was read from channel in
	case v, ok := <-in
		if !ok {
			in = nil //the case will never succeed again
			continue
		}
		
	// process the v that was read from channel in2
	case v, ok := <- in2:
		if !ok {
			in2 = nil
			continue
		}
	case <-done:
		return
	}
}

// How to time out code
// Most interactive programs have to return a response within a certain
// amount of time, with concurrency we can control how much time a request
// has to run (other languages have features on top of promises of futures)
// Go has this functionality with the timeout idiom

func timeLimit() (int, error) {
	// every time you need to limit how long an operation takes in Go,
	// you will see a variation of this pattern
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	select {
	case <-done:
		// we use the goroutine closure to assign values to result and err,
		// and to close the done channel, if this happens first, the read
		// from done succeeds and the values are returned
		return result, err
	case <-time.After(2 * time.Second):
		// the second channel is returned by the After function, it has a value
		// written to it after the specified time.Duration has passed. When this
		// value is read before doSomeWork finiches, timeLimit returns the timeout error
		return 0, errors.New("work timed out")
	}
	// If we exit timeLimit before the goroutine finiches processing, it will continue
	// to run and the result that eventually returns is not used. For stop the work
	// in a goroutine, use context cancellation
}

func main() {

}
