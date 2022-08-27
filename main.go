package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	kanz := person{firstName: "kanz", lastName: "Han"}
	fmt.Println(kanz)
	fmt.Println(kanz.firstName, kanz.lastName)

}
