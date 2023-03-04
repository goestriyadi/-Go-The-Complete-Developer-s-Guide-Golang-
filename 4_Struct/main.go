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
	jim := person{
		firstName: "Jim",
		lastName:  "party",
		contactInfo: contactInfo{
			email:   "jim@gmailcom",
			zipCode: 123,
		},
	}

	jim.updateNmae("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateNmae(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
