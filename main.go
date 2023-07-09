package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jd := person{
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			email:   "john@doe.com",
			zipCode: 12345,
		},
	}

	fmt.Println(jd)
	// Log each value separate by fieldname
	fmt.Printf("%+v", jd)
}
