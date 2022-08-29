package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func (e englishBot) getGreeting() string {
	return "Hello"
}

func (s spanishBot) getGreeting() string {
	return "Hola"
}

func printGretting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGretting(eb)
	printGretting(sb)
}
