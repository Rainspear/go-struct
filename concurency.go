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
	for i := 0; i < len(links); i++ {
		current := time.Now().UnixMilli()
		fmt.Println(<-c)
		fmt.Println("request time", time.Now().UnixMilli()-current)
	}

}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, " is down")
		c <- "Might be down"
		return
	}
	fmt.Println(link, " is up")
	c <- "Might be up"
}
