package main

import (
	"fmt"
	"math/rand"
	"time"
)

type learnFanIn struct {
}

func (learnFanIn) executeMain() {
	c := fanIn(createChannel("Jess"), createChannel("Json"))
	// for v := range c { // will loop forever because input for i := 0; ; i++ create forever
	// 	fmt.Println("Value of those channels: ", v)
	// }
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("About to exit, leaving ...")
}

func createChannel(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) // create string
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string { // make data from many channels into 1 channel
	c := make(chan string)
	go func() {
		for {
			c <- <-input1 // receive from "Jess" when creating channel
		}
	}()
	go func() {
		for {
			c <- <-input2 // receive from "Json" when creating channel
		}
	}()
	return c
}
