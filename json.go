package main

import (
	"encoding/json"
	"fmt"
)

type learnJSON struct{}

type personJSON struct {
	First string
	Last  string
	Age   int
}

func (learnJSON) executeMain() {
	p := personJSON{"Asley", "Ng", 26}

	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
