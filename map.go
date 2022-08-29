package main

import "fmt"

type learnMap struct{}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func (learnMap) executeMain() {
	fmt.Println()
	fmt.Println("Map MODULE")
	colors := map[string]string{"red": "red", "green": "green"}
	colors["white"] = "white"
	delete(colors, "red")
	fmt.Println(colors)
	var newColor map[string]string
	fmt.Println(newColor)
	newColor1 := map[string]string{}
	fmt.Println(newColor1)
	images := make(map[string]string)
	fmt.Println("images", images)
	printMap(colors)
}
