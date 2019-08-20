package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	//fmt.Println(cards.toString())
	cards := newDeckFromFile("my_file")
	cards.shuffle()
	cards.print()
}
