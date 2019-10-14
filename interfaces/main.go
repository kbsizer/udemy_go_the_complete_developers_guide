package main

import "fmt"

// Generig ChatBot interface
type chatBot interface {
	getGreeting() string
}

// PART 1: Using VALUE semantics (method gets copy of object state)

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

func printGreeting(cb chatBot) {
	fmt.Println(cb.getGreeting())
}

// PART 2: Using POINTER (REFERENCE) semantics (object state is shared)
type displayable interface {
	display()
}

type user struct {
	name string
}

// implement "displayable" interface using pointer semantics
func (u *user) display() {
	fmt.Println("Displaying user", u.name)
}

func logUserInfo(d displayable) {
	d.display()
}

func main() {
	fmt.Println("***** INTERFACE DEMO : Part 1 (value semantics) *****")
	eb := englishBot{}
	sb := spanishBot{}

	// polymorphism
	printGreeting(eb)
	printGreeting(sb)

	fmt.Println("***** INTERFACE DEMO : Part 2 (pointer semantics) *****")
	u := user{"Bob"}

	u.display()

	/* The following line will result in compiler error:
	 *     cannot use u (variable of type user) as displayable value
	 *     in argument to logUserInfo: missing method display
	 * because u is passed into logUserInfo() by value and values do
	 * not have access to pointer/reference methods.
	 */
	// logUserInfo(u)
	logUserInfo(&u)
}
