package main

func main() {
	//To create a new deck
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// cards.savetoFile("my_cards")
	// fmt.Println(cards.toString())

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	//Print all the values present in card slice/array
	// cards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
