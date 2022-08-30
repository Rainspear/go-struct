package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

type learnHttp struct{}

func (learnHttp) executeMain() {
	fmt.Println()
	fmt.Println("HTTP MODULE")

	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))
	lw := logWriter{}
	io.Copy(lw, res.Body)
	// io.Copy(os.Stdout, res.Body)
}

func (l logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))
	return len(bs), nil
}
