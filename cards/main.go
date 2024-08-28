package main

func main() {
	cards := newDeck()

	hand, remDeck := deal(cards, 5)

	hand.print()

	remDeck.print()
	// fmt.Println(cards)

}
