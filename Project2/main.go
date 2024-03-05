package main

func main() {
	// Save the cards to local

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
