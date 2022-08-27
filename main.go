package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	// kanz := person{}
	// kanz.firstName = "Kanz"
	// kanz.lastName = "Han"
	// fmt.Printf("%+v\n", kanz)
	// fmt.Println(kanz)
	jim := person{
		firstName: "Jim",
		lastName:  "Robert",
		contact: contact{
			email:   "jimhim@gmail.com",
			zipCode: 40040,
		}}
	jim.print()
	jim.updateName("John")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateName(name string) {
	p.firstName = name
}
