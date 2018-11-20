package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck which is a slice of strings
type deck []string

// d in the function makes it so that you can set up methods for the deck type and can be used on any instance of deck variable
// convention is to put the type name as the receiver, could be anything

// whenever you want to return something from a function, must declare type that will be returned
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// This is a very common pattern
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	// Line below takes a slice of strings and condeses them into 1 string that is comma seperated
	return strings.Join([]string(d), ",")
}

func deal(d deck, handSize int) (deck, deck) {
	// Line below takes a type deck and returns a portion of the slice and the remaining slice
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// When Error handling with Go, Just ask yourself common-sense questions

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// Very common error handling method in GO
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// Line below takes the value of the index and swaps it with a random number for every index in the slice
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
