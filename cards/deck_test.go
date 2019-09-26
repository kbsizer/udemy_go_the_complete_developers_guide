package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert" // $ go get github.com/stretchr/testify
)

// Test_newDeck verifies that newDeck() returns a deck
// 		1) with expected number of cards
// 		2) with expected first card
// 		3) with expected last card
func Test_newDeck_v1(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("newDeck() returned %v cards, want %v cards", len(d), 52)
	}
	firstCard := d[0]
	if firstCard != "Ace of Spades" {
		t.Errorf("newDeck()[0] = %v, want %v cards", firstCard, "Ace of Spades")
	}
	lastCard := d[51]
	if lastCard != "King of Clubs" {
		t.Errorf("newDeck()[0] = %v, want %v cards", lastCard, "King of Clubs")
	}

}

// same as above, but using my homebrew "assert"
func Test_newDeck_v2(t *testing.T) {
	d := newDeck()
	assertEquals(t, "length", len(d), 52)
	assertEquals(t, "first card", d[0], "Ace of Spades")
	assertEquals(t, "last card", d[51], "King of Clubs")
}

func assertEquals(t *testing.T, msg string, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: got %v, wanted %v", msg, got, want)
	}
}

// same as above, but using stretchr/testify's "assert"
func Test_newDeck_v3(t *testing.T) {
	d := newDeck()
	assert.Equal(t, 52, len(d), "length")
	assert.Equal(t, "Ace of Spades", d[0], "first card")
	assert.Equal(t, "King of Clubs", d[51], "last card")

	assert.
}

func Test_deal(t *testing.T) {
	type args struct {
		d        deck
		handSize int
	}
	tests := []struct {
		name  string
		args  args
		want  deck
		want1 deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := deal(tt.args.d, tt.args.handSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("deal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_deck_deal(t *testing.T) {
	type args struct {
		handSize int
	}
	tests := []struct {
		name  string
		d     deck
		args  args
		want  deck
		want1 deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.d.deal(tt.args.handSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deck.deal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("deck.deal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_deck_print(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.print()
		})
	}
}

func Test_deck_toString(t *testing.T) {
	tests := []struct {
		name string
		d    deck
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.toString(); got != tt.want {
				t.Errorf("deck.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deck_saveToFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		d       deck
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.saveToFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("deck.saveToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newDeckFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    deck
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newDeckFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("newDeckFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDeckFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deck_shuffle(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.shuffle()
		})
	}
}
