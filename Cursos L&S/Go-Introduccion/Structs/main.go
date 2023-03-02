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
	//martin := person{firstName: "Martin", lastName: "Gomez"}
	//fmt.Println(martin)
	var alex person
	alex.firstName = "alex"
	alex.lastName = "anderson"
	fmt.Println(alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 1828,
		},
	}
	//jimPointer := &jim
	(&jim).updateName("jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
