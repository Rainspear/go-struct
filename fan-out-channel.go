package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type learnFanOut struct {
}

func (learnFanOut) executeMain() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("program is about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v1 int) {
			c2 <- timeConsumingWork(v1) // simulate side effect
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1e3)
}
