package main

//Static languages like Go, C++, Java you must define data type
// := or manually defining the type on initialization
// only have to put := on initialization of variable
// Go is NOT an Object-orientated language
// Common Go pattern is take a primitive type and extend its functionality to make a new type, like prototype in JS

// Slices are a type of Array that doesnt have a defined length

// type conversion in go is []the-type-you-want(the-value-you-have)
// when doing type conversion think "This is the type I have" and "This is the type I want" and find a way from point A to point B

// Differenct between Receiver and Argument
//// func (p *Page) save() error
//// reads "attach a method called save that returns an error to the type *Page",

//// func save(p *Page) error
//// reads "declare a function called save that takes one parameter of type *Page and returns an error"

func main() {
	// var card string = "Ace of spades"
	cards := newDeck()
	cards.saveToFile("full-deck")
}
