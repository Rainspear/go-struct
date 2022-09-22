package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	total := 1000
	counter := 0
	var wg sync.WaitGroup
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final Value", counter)
}
