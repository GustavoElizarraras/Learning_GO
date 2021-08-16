package main

// Concurrency does not mean more speed, sometimes the algorithms are fast are
// fast enough that the overhead of passing values via concurrency overwhelms 
// any time saved

// Concurrency is used when it is possible to combine data from different and
// independent operations

// Goroutines

// Core concepts
// -Process: instance of a program that runs in an OS with some resources
// -Threads: unit of execution that is given some time to run in the OS
//  with shared resources. One job of an OS is schedule threads
//  on the CPU to make sure every process gets a chance tu run

// When a program starts, the Go runtime creates a number of threads and
// launches a single goroutine to run the program. Some benefits are:
// -Goroutine creation is faster than thread creation
// -Goroutine initial stack is smaller and scalable than threads ones
// -Switching between goroutines is faster than switching between threads
// -The scheduler is able to optimize its decisions
// All of these allow to spawn ten of thousand of goroutines

// A go routine is launched by placing the "go" keyword before (any) function
// invocation, one thing to consider is that any values returned are ignored

// It is customary in Go to launch goroutines with a closure that wraps 
// business logic, it takes care of the current bookkeeping.

func process(val int) {
	// do domething with val
}

func runThingConcurrently(in <-chan int, out chan<- int) {
	go func() {
		for val := range in { // read values out of channels
			result := process(val) // passes them to the business logic
			out <- result // result of the function is passed to a diff channel
		}
	}() // closure does not know it is inside a goroutune
}

// Channels
// Goroutines communicate using channels, built-in type created using make()
ch := make(chan int)
// If a channel is passed to a function, you are passing a pointer to the channel
// and its zero value is nil

// Reading, Writing and Buffering
// <- operator to interact with a channel:
a := <- ch // reads a value from ch, assigns to a
ch <- b // write the value in b to ch
// Each value wrtitten to a channel can only be read once. If multiple goroutines are
// reading from the same channel, a value written to the channel will only be read by one 

// It is rare for a goroutine to read and write to the same channel. When assigning
// to a variable or field: (ch <-chan int) to indicate that the goroutine only reads 
// to indicate it only writes: (ch chan<- int)

// By default channels are unbuffered. Every write to an open, unbuffered channel causes the 
// writing goroutine to pause until another goroutine reads from the same channel and vice versa.
// So, you cannot write to or read from an unbuffered channel without at least two concurrently
// running goroutines

// Go also has buffered channels, these channels buffer a limited number of writes without blocking.
// If the buffer fills before any reads from the channel, a subsequent write to the channel pauses the 
// goroutine until the channel is read; reading from a channel with an empty buffer also blocks.
ch := make(chan int, 10) // buffered channel, specifying the capacity. cap() and len() work too

// Read from a channel using a for-range loop:
for v := range ch { // only one variable declared 
	fmt.Println(v)
	// It stops by a break, return or if the channel is closed
}

// Closing a channel 
close(ch)
// Writes panic
// Reading succeeds if the channel has buffered values, if not, the zero value for the channel type
// is returned

// We use the comma ok idiom to detect whether a channel has been closed or not
v, ok := <-ch // if ok == true, the channel is open

// The responsability for closing a channel lies with the goroutine that writes to the channel. It
// only need to be close if there is a goroutine waiting for the channel to close.

// Channels set apart Go's concurrency model. They guide us into thinking about our code as a series of 
// stages and making data dependencies clear 


func main() {

}
