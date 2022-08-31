package main

import (
	"fmt"
	"net/http"
)

type learnConcurrency struct{}

func (lc learnConcurrency) executeMain() {
	fmt.Println()
	fmt.Println("Concurrency MODULE")

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"https://go.dev/",
	}

	c := make(chan string)

	for _, v := range links {
		go checkLink(v, c)
	}

	for l := range c {
		// remember to add argument because anonymous function will create a new memory for variable
		// and not update to new value of channel so that
		// go func() {
		// 	time.Sleep(5 * time.Second)
		// 	checkLink(l, c) // warning because not update new value of "channel" and still use old value from channel variable
		// }()
		// go func(link string) {
		// 	time.Sleep(5 * time.Second)
		// 	checkLink(link, c)
		// }(l)
		fmt.Println(l)
	}
	// for {
	// 	fmt.Println(<-c)
	// }

	// This will cause error: because all goroutines are asleep - deadlock! => no consumption value in channel
	// c1:= make(chan int)
	// fmt.Println(<-c1)

	// This will run normally maybe channel now is a buffer
	// c1:= make(chan int, 1)
	// fmt.Println(<-c1)

	// This will run normally because it will separate in another thread so main thread not error
	// c1:= make(chan int)
	// go func() {
	// c <- 3
	// }()
	// fmt.Println(<-c1)
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, " is down")
		c <- link
		return
	}
	fmt.Println(link, " is up")
	c <- link
}
