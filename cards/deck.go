package main

import "fmt"

// Create a new type: Deck
// Adds type safety and allows us to add methods to our "slice of strings"
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}