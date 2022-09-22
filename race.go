// Program with race condition
package main

import (
	"fmt"
	"runtime"
	"sync"
)

type learnRace struct{}

var GFG = 0

// This is the function we’ll run in every
// goroutine. Note that a WaitGroup must
// be passed to functions by pointer.
func worker(wg *sync.WaitGroup) {
	GFG = GFG + 1

	// On return, notify the
	// WaitGroup that we’re done.
	wg.Done()
}
func (learnRace) executeMain() {

	// This WaitGroup is used to wait for
	// all the goroutines launched here to finish.
	var w sync.WaitGroup

	// Launch several goroutines and increment
	// the WaitGroup counter for each
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go worker(&w)
	}
	// Block until the WaitGroup counter
	// goes back to 0; all the workers
	// notified they’re done.
	fmt.Println("GoRoutine:", runtime.NumGoroutine())
	w.Wait()
	fmt.Println("Value of x", GFG)
}
