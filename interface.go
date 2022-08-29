package main

import "fmt"

type learnInterface struct{}

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

func (learnInterface) executeMain() {
	fmt.Println()
	fmt.Println("Interface MODULE")
	eb := englishBot{}
	sb := spanishBot{}
	printGretting(eb)
	printGretting(sb)
}
