package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cards = addCardsToDeck()

	return cards
}

func addCardsToDeck() deck {
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
