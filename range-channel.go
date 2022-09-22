package main

import "fmt"

type learnRangeChannel struct{}

func (learnRangeChannel) executeMain() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) // close channel when finish transfer data -> not close will be deadlocked when using range
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("End of channel")
}
