package main

import "fmt"

type learnQuitChannel struct{}

func (learnQuitChannel) executeMain() {
	e := make(chan int)
	o := make(chan int)
	q := make(chan int)

	go sendMessage(e, o, q)
	receiveMessage(e, o, q)

}

func sendMessage(e, o, q chan int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receiveMessage(e, o, q chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Even Number: ", v)
		case v := <-o:
			fmt.Println("Odd Number: ", v)
		case v := <-q:
			fmt.Println("Quit Channel here ", v)
			return
		}
	}
}
