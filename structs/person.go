package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func NewPerson(firstName, lastName string, contactInfo ContactInfo) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		contact:   contactInfo,
	}
}

func (p Person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p *Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
