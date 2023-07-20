package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Create a New Type of 'deck'
//which is 	a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamond", "Heart", "Clubs"}
	cardValue := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	// []string(d)
}

func (d deck) savetoFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
