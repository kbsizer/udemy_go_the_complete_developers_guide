package main

import "fmt"

// Generig ChatBot interface
type chatBot interface {
	getGreeting() string
}

// Concrete English Bot implementation
type englishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello"
}

// Concrete Spanish Bot implementation
type spanishBot struct{}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func main() {
	fmt.Println("* INTERFACE DEMO *")
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(cb chatBot) {
	fmt.Println(cb.getGreeting())
}
