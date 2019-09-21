package main

// reversibleString is a simple example of a user-defined type
// and method in Go.  Also illustrates the correct idium for
// working with characters and runes.
type reversibleString string

func (s reversibleString) Reverse() string {
	runes := []rune(string(s))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
