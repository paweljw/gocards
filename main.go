package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cards := newDeck()
	cards.print()

	fmt.Println("\nAnd the shuffled hand is:")
	hand, cards := deal(cards, 5)
	hand.shuffle()
	hand.print()

	fmt.Println("Deck size is: ", len(cards))

	fmt.Println(cards.toString())

	_ = cards.saveToFile("deck.txt")

	fmt.Println("+++")
	newDeck := newDeckFromFile("deck.txt")
	newDeck.print()
}
