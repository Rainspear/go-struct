package main

import (
	"fmt"
	"sync"
	"time"
)

type learnSelect struct {
}

func (learnSelect) executeMain() {
	timeSend := 10
	var wg sync.WaitGroup
	wg.Add(2)
	go loopWhenSend(timeSend, &wg)
	go loopwhenReceive(timeSend, &wg)
	wg.Wait()
}

func loopWhenSend(timeSend int, wg *sync.WaitGroup) {

	x, y := 0, 1
	c := make(chan int)
	d := make(chan int)
	go func() {
		for i := 0; i < timeSend; i++ {
			select {
			case c <- x:
				time.Sleep(time.Millisecond * 200)
				x++
			case d <- y:
				y += x
				time.Sleep(time.Millisecond * 200)
			}
		}
		close(c)
		close(d)
	}()

	// print 10 times
	fmt.Println("loopWhenSend x:", <-c)
	fmt.Println("loopWhenSend y:", <-d)
	fmt.Println("loopWhenSend y:", <-d)
	fmt.Println("loopWhenSend x:", <-c)
	fmt.Println("loopWhenSend y:", <-d)
	fmt.Println("loopWhenSend x:", <-c)
	fmt.Println("loopWhenSend y:", <-d)
	fmt.Println("loopWhenSend y:", <-d)
	fmt.Println("loopWhenSend y:", <-d)
	fmt.Println("loopWhenSend x:", <-c)
	// quit <- 0
	wg.Done()
}

func loopwhenReceive(timeSend int, wg *sync.WaitGroup) {
	x, y := 0, 1
	c := make(chan int)
	d := make(chan int)

	go func() {
		for i := 0; i < timeSend; i++ {
			select {
			case c <- x:
				time.Sleep(time.Millisecond * 200)
				x++
			case d <- y:
				y += x
				time.Sleep(time.Millisecond * 200)
			}
		}
		close(c)
		close(d)
	}()

	for i := 0; i < timeSend/2; i++ {
		fmt.Println("loopwhenReceive x:", <-c)
		fmt.Println("loopwhenReceive y:", <-d)
	}

	wg.Done()
}
