package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	total := 1000
	t := time.Now()
	var counter int64
	var wg sync.WaitGroup
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Final Value", counter)
	fmt.Println("Final Time", time.Since(t))
}
