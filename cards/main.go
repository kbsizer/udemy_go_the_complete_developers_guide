package main

import "fmt"

const deckSize = 52

func main() {
	// var card string = "Ace of Spades"
	card := newCard()
	fmt.Println(card)
	fmt.Println(deckSize)
}

func newCard() string {
	return "Five of Diamonds"
}

func estPi() float64 {
	return 3.14
}
