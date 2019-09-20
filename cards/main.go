package main

import "fmt"

const deckSize = 52

func main() {
	// var card string = "Ace of Spades"
	card := newCard()
	fmt.Println(card)
	fmt.Println(deckSize)

	// a slice of type string
	cards := []string{"Ace of Spades", "Jack of Diamonds", "3 of Clubs"}
	fmt.Println("cards =", cards)
	fmt.Printf("cards = %v\n", cards)
	// LESSON 16

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

//my attempt at a method
type reversibleString string

func (s reversibleString) Reverse() string {
	runes := []rune(string(s))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
