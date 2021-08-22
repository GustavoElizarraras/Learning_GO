package main

import "sync"

// When to use mutexes instead of channels
// mutexes are used to coordinate access to data across threads, they also
// limit the concurrent execution of some code. BUt they usually obscure
// the data flow of the program. THe common case is when your goroutines
// read or write a shared value, but don't process it.

func scoreboardManager(in <-chan func(map[string]int), done <-chan struct{}) {
	// listens on one channel for a function that reads or modifies the map on a
	// second channel to know wheen to shut down
	scoreboard := map[string]int{} // map declaration
	for {
		select {
		case <-done:
			return
		case f := <-in:
			f(scoreboard)
		}
	}
}

type ChannelScoreboardManager chan func(map[string]int)

func NewChannelScoreboardManager() (ChannelScoreboardManager, func()) {
	ch := make(ChannelScoreboardManager)
	done := make(chan struct{})
	go scoreboardManager(ch, done)
	return ch, func() {
		close(done)
	}
}

func (csm ChannelScoreboardManager) Update(name string, val int) {
	// puts a value into the map
	csm <- func(m map[string]int) {
		m[name] = val
	}
}

// What about reading from the scoreboard? We need to return a value back
func (csm ChannelScoreboardManager) Read(name string) (int, bool) {
	var out int
	var ok bool
	done := make(chan struct{})
	csm <- func(m map[string]int) {
		out, ok = m[name]
		close(done)
	}
	<-done
	return out, ok
}

// All the previous code works, but only allows a single reader at a time

// Using mutexes

// Implementation 1: Mutex has two methods, Lock and Unlock. Lock causes the
// current goroutine to pause as long as another goroutine is currently in the
// critical section. When it is cleared, the lock is acquired by the current
// goroutine and the code in the critical section is executed; and Unlock marks
// the end of the critical section

// Implementation 2: RWMutex, allos to have reader and writer locks. Reader locks
// are shared, multiple readers can be in the critical section at once. The writer
// lock is manged with Lock and Unlock, meanwhile the reader lock with RLock and RUnlock

// Anytime a mutex lock is acquired, we must make sure it is released, it can be
// done with a defer statement to call Unlock

// Tips:
// If a goroutine tries to acquire the same lock twice (or 2 differentfunctions), it
// deadlocks and waits for itself to release the lock. If a function is recursive, make
// sure to release it before it is called from within.

// Never try to access a variable from multiple gorouties unless you acquire a mutex for
// that variable first, it can cause odd errors that are hard to find

// In rare situations where you need to share a map across multiple goroutines, use a
// built-in map protexted by a sync.RWMutex. sync.Map is used where key/value pairs are
// inserted once and read many times and when goroutines share the map, but don't access
// each other's keys and values

type MutexScoreboardManager struct {
	l          sync.RWMutex
	scoreboard map[string]int
	// there is no transfer of scoreboard, using a mutex makes sense.
	// Also it is stored in memory, if it is in a server or db, don't use
	// mutexes t guard access to the system
}

func NewMutexScoreboardManager() *MutexScoreboardManager {
	return &MutexScoreboardManager{
		scoreboard: map[string]int{},
	}
}

func (msm *MutexScoreboardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = val
}

func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
	msm.l.Lock()
	defer msm.l.Unlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}

// Considerations:

// If you are coordinating goroutines or tracking a value as it is transformed
// by a series of goroutines, use channels

// If you are sharing access to a field in a struct, use mutexes

// If channels are slow, and you cant get around it, use mutexes

func main() {

}
