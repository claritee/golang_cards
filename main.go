package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Sizx of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
