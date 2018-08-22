package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	// for i := 0; i < 15; i++ {
	// 	fmt.Println(rand.Intn(3))
	// }

	// save(toString(cards), "deck")
	// cards := read("deck")
	fmt.Println("Before Shuffle")
	// cards.print()
	// fmt.Println("After Shuffle")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "King of Diamonds"
}
