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

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}
