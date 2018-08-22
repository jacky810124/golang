package main

import (
	"fmt"
)

func main() {
	cards := []string{"Ace Of Diamonds", newCard()}
	cards = append(cards, "2222")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "King of Diamonds"
}
