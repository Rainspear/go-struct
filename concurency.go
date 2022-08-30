package main

import (
	"fmt"
	"net/http"
	"time"
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
		current := time.Now().UnixMilli()
		fmt.Println(l)
		fmt.Println("request time", time.Now().UnixMilli()-current)
		go checkLink(l, c)
	}
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
