package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Using a cancel function to terminate a goroutine
// We can use the done channel pattern to return a cancelation alongside.
func countTo(max int) (<-chan int, func()) {
	// Two channels, one that returns data and another for signaling done
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		// CLosure that closes the done channel and return the closure
		close(done)
		// Cancelling with a closure allows us to perform additional clean up work
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			default:
				ch <- i
			}
		}
		close(ch)
	}()
	return ch, cancel
}

// Buffered vs Unbuffered channels
// By default, channels are Unbuffered, they are easy, one goroutine writes
// and waits for another goroutine to pick up its work. Buffered channels are
// much more complicated, you have to pick a size, and the case where the
// channel is full and the writting goroutine blocks waiting for a reading
// goroutine, must be handled

// Buffered channels are useful when you know how manu goroutines you have launched,
// want to limit the number of goroutines you will launch or limit the amount of work
// that is queued up. They work great for gathering data from a set of gouroutines
// or limit concurrent usage, also, for managing the system workload.

// func processChannel(ch chan int) []int {
// 	const conc = 10
// 	results := make(chan int, conc)
// 	for i := 0; i < conc; i++ {
// 		// processing the first 10 results on a channel, todo this, we
// 		// launch 10 goroutines, each writes its results to a buffered channel
// 		go func() {
// 			results <- process(v)
// 		}()
// 		// We know how many goroutines have launched, and each of them exit as soon
// 		// as it finishes its work. The buffered channel has one space for each
// 		// launched goroutine without blocking.
// 	}
// 	var out []int
// 	for i := 0; i < conc; i++ {
// 		// Looping over the buffered channel reading out the values as they are written,
// 		// when all the values have been read, we return the results knowing we are not
// 		// leaking any goroutine
// 		out = append(out, <-results)
// 	}
// 	return out
// }

// Backpressure
// A technique for buffered channels, it is counterintuitive but systems perform better
// overall because the work amount is limited.
// We can use a buffered channel and a select statement to limit the number of
// simultaneous requests ina a system :

type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	// struct that contains a buffered channels with a number of tokens
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

// function to run
func (pg *PressureGauge) Process(f func()) error {
	select {
	// tries to read a token from the channel, if it can, function runs,
	// if not, default that raises an error
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}

func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func main() {

	ch, cancel := countTo(10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	cancel()

	// Real implementation of backtraking, when more than 2 requests are
	// asked, it returns no more capacity
	pg := New(2)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)

}
