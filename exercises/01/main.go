package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func even(i int) {
	if i%2 == 0 {
		fmt.Println("Even:", i)
	}
	wg.Done()
}

func odd(i int) {
	if i%2 != 0 {
		fmt.Println("Odd:", i)
	}
}

func main() {

	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go even(i)
	// }
	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go odd(i)
	// }

	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	wg.Add(2)
	go func() {
		fmt.Println("routine1")
		wg.Done()
	}()
	go func() {
		fmt.Println("routine2")
		wg.Done()
	}()
	fmt.Println("Executing")
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}
