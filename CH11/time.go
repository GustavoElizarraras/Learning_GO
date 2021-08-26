package main

import (
	"fmt"
	"time"
)

// There are two main types used to represent time:

// -*- time.Duration, period of time, int64, the smallest is 1 ns
// 	 also, strings are supported with the time.ParseDuration(), 2h45m
//   There are methods on time.Duration to extract portions of it, like
//   Day, Month, Minute, etc.

// -*- time.Time, a moment of time, with a time zone. time.Now
//	  Compare 2 time.Time instances with After, Before and Equal methods
//   The Sub mehod returns a time.Duration difference between two time.Time
//   There are a lot more methods, check the documentation

// Monotonic Time
// OS's keep track of two different sorts of time: wall clock (current time)
// and monotonic clock (count up from the time the computer was booted)
// The wall clock does not uniformly increase, so this could cause problems,
// to address this problem, Go uses monotonic time when a timer is set or with
// or a time.Time instance is created with time.Now

// Timers and Timeouts
// The time package includes functions that return channels that output values
// after a specified time. time.After returns a channel that outputs once,
// meanwhile time.Tick returns a new value every time the specified time.Duration
// elapses, these are used with Go's concurrency support to enable timeouts or
// recurring tasks. There are more functions, like time.AfterFunc, for more
// information, check the documentation

func main() {
	d := 2*time.Hour + 30*time.Minute // d is of type time.Duration
	fmt.Println(d)
	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2016-13-03 00:00:00 +0000")
	// converts from a string to a time
	if err != nil {
		return
	}
	fmt.Println(t, t.Format("January 2, 2006 at 3:04:05PM MST"))
	// converts a time.Time to a string
}
