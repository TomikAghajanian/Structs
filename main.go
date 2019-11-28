package main

import (
	"fmt"
)

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
		lastName:  "Party",
		contactInfo: contactInfo {
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Tom")
	jim.print()
}

func (pointerToPerson *person) print(){
	fmt.Printf("%+v", *pointerToPerson)
}
func (pointerToPerson *person)  updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}