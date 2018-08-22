package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"♠︎", "♥︎", "♦︎", "♣︎"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := suit + value

			cards = append(cards, card)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func toString(d deck) string {
	return strings.Join(d, " ")
}

func save(s string, filename string) error {
	return ioutil.WriteFile(filename, []byte(s), 0666)
}

func read(filename string) deck {
	result, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(result), " "))
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := range d {
		targetIndex := rand.Intn(len(d))

		d[i], d[targetIndex] = d[targetIndex], d[i]
	}
}
