package main

func main() {
	cards := newDeckFromFile("mycards")

	cards.print()

	cards.shuffle()

	cards.print()

	// cards.saveToFile("mycards")

	// hand, remainingDeck := deal(cards, 5)

	// remainingDeck.print()

	// hand.print()
}
