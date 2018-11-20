package main

import (
	"fmt"
)

// Struct is like an object consructor from javascript

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

// %+v shows the property name as we as the value
func main() {
	jarrod := person{
		firstName: "Jarrod",
		lastName:  "Mcatee",
		contactInfo: contactInfo{
			email: "jarrod@gmail.com",
			zip:   28273,
		}}
	// & is an operater for a variable which says give me the address in the RAM that the following variable is pointing at
	jarrodPointer := &jarrod
	jarrodPointer.updateFirstName("Jarrodd")
	jarrod.print()

	// if you have a value, you can turn it into a address with the & operater, see line 29
	// if you have a address, you can get the value of that address with the * operater, see line 40
}

// * operater in front of a type, it means that its a description of a type, making it so that the function can only be called on that type
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	// * operater when before a pointer is saying give me access to the information at that address
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
}
