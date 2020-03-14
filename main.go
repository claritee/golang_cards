package main

func main() {
	// call function
	// cards := newDeck()

	// looping
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// calling instance function
	//cards.print()

	// assigning multiple return values
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// type conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	cards := newDeck()
	// fmt.Println(cards.toString())
	cards.saveTofile("my_cards")
}

func newCard() string {
	return "Five of Diamonds"
}
