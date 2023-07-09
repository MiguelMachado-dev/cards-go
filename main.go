package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var jd person

	jd.firstName = "John"
	jd.lastName = "Doe"

	fmt.Println(jd)
	// Log each value separate by fieldname
	fmt.Printf("%+v", jd)
}
