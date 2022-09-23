package main

import (
	"fmt"
	"time"
)

type learnSelect struct {
}

func (learnSelect) executeMain() {
	timeSend := 10

	loopWhenSend(timeSend)
	loopwhenReceive(timeSend)
}

func loopWhenSend(timeSend int) {
	x, y := 0, 1
	c := make(chan int)
	d := make(chan int)
	go func() {
		for i := 0; i < timeSend; i++ {
			select {
			case c <- x:
				time.Sleep(time.Second)
				x++
			case d <- y:
				y += x
				time.Sleep(time.Second)
			}
		}
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

	// for i := 0; i < timeSend/2; i++ {
	// 	fmt.Println("x:", <-c)
	// 	fmt.Println("y:", <-d)
	// }
}

func loopwhenReceive(timeSend int) {
	x, y := 0, 1
	c := make(chan int)
	d := make(chan int)
	for i := 0; i < timeSend; i++ {
		select {
		case c <- x:
			time.Sleep(time.Second)
			x++
		case d <- y:
			y += x
			time.Sleep(time.Second)
		}
	}

	for i := 0; i < timeSend/2; i++ {
		fmt.Println("loopwhenReceive x:", <-c)
		fmt.Println("loopwhenReceive y:", <-d)
	}
}
