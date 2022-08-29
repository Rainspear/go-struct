package main

import (
	"fmt"
	"net/http"
	"os"
)

type learnHttp struct{}

func (learnHttp) executeMain() {
	fmt.Println()
	fmt.Println("HTTP MODULE")

	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println(res)
}
