package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	jd := person{firstName: "John", lastName: "Doe"}

	fmt.Println(jd)
}
