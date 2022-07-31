package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cards = addCards()

	return cards
}

func addCards() deck {
	cardSuits := []string{"Ace", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	deck := deck{}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			deck = append(deck, cardSuit+" of "+cardValue)
		}
	}

	return deck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(card, i)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	const DefaultPermissionsForWriteAndRead = 0666
	return ioutil.WriteFile(filename, []byte(d.toString()), DefaultPermissionsForWriteAndRead)
}

func newDeckFromFile(filename string) deck {
	deckBs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	stringSlice := convertByteSliceToString(deckBs)
	return deck(stringSlice)
}

func convertByteSliceToString(deckBs []byte) []string {
	return strings.Split(string(deckBs), ",")
}
