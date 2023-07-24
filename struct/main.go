package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int32
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// fmt.Println("Welcome to Main")
	// alex := person{firstName: "Alex", lastName: "Alexander"}
	// fmt.Println(alex)

	// var alex person
	// alex = {firsName: "alex"}
	// alex.firstName = "Priyansh"
	// alex.lastName = "Parikh"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jimmy",
		lastName:  "Fellon",
		contact: contactInfo{
			email:   "jimmyxyz@gmail.com",
			zipCode: 560095,
		},
	}

	// fmt.Printf("%+v", jim)
	// jimpointer := &jim
	jim.updateName("Priyansh")
	jim.print()
}

func (pointertoPerson *person) updateName(newFirstName string) {
	(*pointertoPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
