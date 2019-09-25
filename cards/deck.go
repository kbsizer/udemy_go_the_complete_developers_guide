package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// deck defineds a new type
//
// Note: This adds type safety and allows us to add methods to our "slice of strings".
type deck []string

// newDeck creates and returns a new deck.
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// deal FUNCTION splits the given deck into a hand of handSize and the remaining cards.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// deal METHOD
func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// print writes the contents of this deck to stdio.
//
// Note: this method uses value semantics (it works with a private copy of the deck).
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
	fmt.Println("----------------")
}

// delimiterChar is the character used to delimit individual cards when persisting a deck
// to a file.
const delimiterChar = ","

// toString returns this deck as a comma-delimited string using strings.Join().
func (d deck) toString() string {
	return strings.Join([]string(d), delimiterChar)
}

// save converts this deck to a byte slice and writes it to a file with the given name.
func (d deck) saveToFile(filename string) error {
	s := d.toString()
	deckAsByteSlice := []byte(s)
	return ioutil.WriteFile(filename, deckAsByteSlice, 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	deckAsByteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err) // best practice for logging in Go??
		return nil, err
	}
	s := string(deckAsByteSlice)
	return deck(strings.Split(s, delimiterChar)), err
}
