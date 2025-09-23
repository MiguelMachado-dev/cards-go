package main

func main() {
	// cards := newDeck()

	// cards.saveToFile("my_card")

	cards := newDeckFromFile("my_cards")

	cards.print()
}
