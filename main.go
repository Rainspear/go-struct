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
	/*
		jimPointer := &jim
		jimPointer.updateName("John")
		two line above are the same as jim.updateName("John") and can be used correctly
		because go automatically convert Person into *Person type
		struct, string, int, float, boolean are value type (primitive types)
	*/
	jim.print() // Jim
	jim.updateName("John")
	jim.print() // John

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	updateNumber(numbers)
	fmt.Println(numbers) // [20 2 3 4 5 6 7 8 9 10] because slices, channels, maps, pointers, functions are reference type

	name := "bill"
	namePointer := &name
	fmt.Println("name: ", &name)              // 0xc00005c250
	fmt.Println("namePointer: ", namePointer) // 0xc00005c250
	fmt.Println("&namePointer", &namePointer) // 0xc00000a048
	printPointer(namePointer)
	/*
		printPointer namePointer 0xc00005c250
		printPointer *namePointer bill
		printPointer &namePointer 0xc00000a050
	*/
}

func updateNumber(nums []int) {
	nums[0] = 20
}

func (p *person) print() {
	fmt.Printf("%+v %p \n", *p, &p)
}

func (pointPerson *person) updateName(name string) {
	(*pointPerson).firstName = name
	(*pointPerson).print()
}

func printPointer(namePointer *string) {
	fmt.Println("printPointer namePointer", namePointer)
	fmt.Println("printPointer *namePointer", *namePointer)
	fmt.Println("printPointer &namePointer", &namePointer)
}
