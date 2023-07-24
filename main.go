package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jd := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "john@doe.com",
			zipCode: 12345,
		},
	}

	jd.updateName("Jane")
	jd.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
