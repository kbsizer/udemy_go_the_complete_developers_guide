package main

import (
	"fmt"
	"math" // see: https://golang.org/pkg/math/
)

const deckSize = 52

func main() {
	// var card string = "Ace of Spades"
	card := newCard()
	fmt.Println(card)
	fmt.Println(deckSize)

	// a slice of type string
	cards := deck{"Ace of Spades", "Jack of Diamonds", "3 of Clubs"}
	fmt.Println("cards =", cards)
	fmt.Printf("cards = %v\n", cards)
	// LESSON 16

	cards2 := append(cards, "Queen of Hearts")
	fmt.Println("cards  =", cards)
	fmt.Println("cards2 =", cards2)

	cards.print()

	fmt.Printf("\n====================================\n\n")

	r := reversibleString("1234567890 The quick brown 狐 jumped over the lazy 犬")
	fmt.Println("r           = ", r)
	fmt.Println("r.Reverse() = ", r.Reverse())
}

func newCard() string {
	return "Five of Diamonds"
}

func estPi() float64 {
	return 3.14
}

func hypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}
