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

	s := `[{"First":"Asley","Last":"Ng","Age":26}, {"First":"Asley","Last":"Ng","Age":26}]`
	var p1 []personJSON
	bs1 := []byte(s)
	json.Unmarshal(bs1, &p1)

	fmt.Println("All data of people", p1)
	for k, v := range p1 {
		fmt.Printf("index %d: %+v \r\n", k, v)
	}
}
