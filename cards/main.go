package main

import (
	"fmt"
	"log"
	"math" // see: https://golang.org/pkg/math/
	"reflect"
)

const deckSize = 52

func main() {
	fmt.Println("\n====== illustrates type conversion (string <==> slice of byte) ======")

	s1 := "Simple ASCII string"
	fmt.Println("s1 =", s1)
	b1 := []byte(s1)
	fmt.Println("[]byte(s1) =", b1)
	fmt.Println("and back:", string(b1))
	s2 := "More complex ހ  	ސ  	ޤ  	ހި string with multibyte 𪜀 𪮘 𪾀 𫜴 characters Ꙃ ꙉ ꙮ Ꚗ"
	fmt.Println("s2 =", s2)
	b2 := []byte(s2)
	fmt.Println("[]byte(s2) =", b2)
	fmt.Println("and back:", string(b2))

	// ---------------------------- this whole block would make a good unit test ------
	// create a deck of cards
	cards := newDeck()
	fmt.Println("cards.toString() =", cards.toString())
	// save deck to disk
	cards.saveToFile("myCards.dat")
	// load it back from disk
	cards2, err := newDeckFromFile("myCards.dat")
	if err != nil {
		log.Fatal("Failure while loading from file:", err)
	}
	// compare them
	fmt.Println("Cards from disk same as cards saved to disk?", reflect.DeepEqual(cards, cards2))

	// ----------------------------

	fmt.Println("\n====== shuffle our deck ======")
	cards.shuffle()
	cards.print()

	fmt.Println("\n====== illustrates append() behavior; the input args are not mutated ======")

	// cardsMutated := append(cards, "Queen of Hearts")
	// cards.print()
	// cardsMutated.print()

	// fmt.Println("\n====== deal(): FUNCTION vs METHOD ======")

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// fmt.Printf("\n====================================\n\n")

	// r := reversibleString("1234567890 The quick brown 狐 jumped over the lazy 犬")
	// fmt.Println("r           = ", r)
	// fmt.Println("r.Reverse() = ", r.Reverse())
}

func estPi() float64 {
	return 3.14
}

func hypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}
