package main

import (
	"fmt"
)

// select
// This statement is the control structure for concurrency in Go, it solves
// a common problem: if you can perform two concurrent operations, which one
// do you do first?. You can't favor one operation over others, because maybe
// you will never process some cases, this is called starvation

select {
// each case in a select is a read or write to a channel, if it es possible for
// a case, it is executed along with the body of the case. If multiple cases are
// true, select selects one randomly that can go forward.

case v := <-ch:
	fmt.Println(v)
case v := <-ch2:
	fmt.Println(v)
case ch3 <- x:
	fmt.Println("wrote",x)
case ch4 <-:
	fmt.Println("got value on ch4, but ignored it")
}

// Another advantage of select is acquring locks in inconsistent order. If you have
// two goroutines that access the same two channels, they must be accessed in the
// same order in both, if not, this causes a deadlock; it means neither can proceed
// because they are waiting on each other. If every goroutine in your Go app is
// deadlocked, the Go runtime kills your program

func ex1() {
	// example with deadlock
	// The goroutine that we launch cannot proceed until ch1 is read,
	// and the main gorotuine cannot proceed until ch2 is read
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	ch2 <- v
	v2 := <-ch1
	fmt.Println(v, v2)
}

func ex2() {
	// wrapping the channel accesses in the main goroutine in a select
	// avoid the deadlock
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		// this goroutine writes the value 1 into ch1, so the read from
		// ch1 into v2 in the main goroutine is able to succeed
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	var v2 int
	select {
	// checks if any of its cases can proceed, the deadlock is avoided
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(ch1, ch2, v, v2)
}

// select is responsible for communicating over a number of channels
// is often embedded within a for loop
for {
	select{
	case <- done
	// you must include a way to exit the loop 
		return
	case v := <-ch:
		fmt.Println(v)
	default:
	// Is selected when there are no cases with channels that can be read
	// or written. Is used to implement a nonblocking r/w on a channel
		fmt.Println("no value written in any channel")
	}
}

func main() {

}
