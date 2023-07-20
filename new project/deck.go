package main

import "fmt"

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
	for _, card := range d {
		fmt.Println(card)
	}
}
