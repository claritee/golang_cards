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

	// save to file
	//cards := newDeck()
	// fmt.Println(cards.toString())
	//cards.saveTofile("my_cards")

	// read from file
	cards := newDeckFromFile("my_cards")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
