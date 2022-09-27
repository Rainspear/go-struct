package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

type learnContext struct {
}

func (learnContext) executeMain() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error check 1: ", ctx.Err())
	fmt.Println("Num 1 goroutines: ", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working: ", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 2: ", ctx.Err())
	fmt.Println("Num 2 goroutines: ", runtime.NumGoroutine())
	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 3: ", ctx.Err())
	fmt.Println("Num 3 goroutines: ", runtime.NumGoroutine())

}

/*
	Concept to use context to cancel request or something processing need to be cancel
	1. create context
	ctx := context.Background() // for testing purpose
	ctx := context.TODO() // for uncertain job

*/
