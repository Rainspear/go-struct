package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

type learnGroup struct{}

func (learnGroup) executeMain() {
	t := time.Now()
	printCPU()
	wg.Add(1) // increase 1
	go foo()

	bar()
	printCPU()
	wg.Wait() // wait until wg counter is zero
	fmt.Println(time.Since(t))
}

func printCPU() {
	fmt.Println("OS \t", runtime.GOOS)
	fmt.Println("ARCH \t", runtime.GOARCH)
	fmt.Println("CPU \t", runtime.NumCPU())
	fmt.Println("Goroutines \t", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo", i)
	}
	wg.Done() // decrease 1
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar", i)
	}
}
