package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	total := 1000
	t := time.Now()
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Final Value", counter)
	fmt.Println("Final Time", time.Since(t))
}
