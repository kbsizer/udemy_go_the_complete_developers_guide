package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type personv2 struct {
	firstName   string
	lastName    string
	contactInfo // shorthand; field name same as type
}

func main() {
	fmt.Println("Hello from structs/main.go")
	// struct initialization using positional args
	alex := person{"Alex", "Anderson", contactInfo{"aanderson@gmail.com", 24549}}
	fmt.Println("alex =", alex)

	// struct initialization using named args
	bob := person{
		firstName: "Bob",
		lastName:  "Miller",
		contact: contactInfo{
			email:   "bob@nowhere.com",
			zipCode: 12345,
		},
	}
	fmt.Println("bob =", bob)

	// struct initialization with zero values
	// and using $+v to see field names
	var billy person
	billy.logDetails("empty billy")
	billy.firstName = "Billy Bob"
	billy.lastName = "Thorton"
	billy.logDetails("billy")

	// demo of pass-by-value vs pass-by-reference method semantics
	alex.logDetails("before updateNamePBV ")
	alex.updateNamePBV("CHANGED")
	alex.logDetails(" after updateNamePBV ")
	(&alex).updateNamePBR1("CHANGED1")
	alex.logDetails(" after updateNamePBR1")
	alex.updateNamePBR2("CHANGED2")
	alex.logDetails(" after updateNamePBR2")

	// Go Gotcha: parameters are passed by value BUT elements of slices/arrays ARE modifiable
	mySlice := []string{"A", "B", "C", "D", "E"}
	fmt.Println("mySlice BEFORE =", mySlice)
	updateSlice(mySlice)
	fmt.Println("mySlice AFTER  =", mySlice)
}

func updateSlice(s []string) {
	s[1] = "X"
}

func (p person) logDetails(msg string) {
	fmt.Printf("%s = %+v\n", msg, p)
}

// pass-by-value (method operates on its own private copy)
func (p person) updateNamePBV(newName string) {
	p.firstName = newName
}

// pass-by-reference using conventional C syntax
// (updates the object on which the method is called)
func (p *person) updateNamePBR1(newName string) {
	(*p).firstName = newName
}

// pass-by-reference using GO's shorthand
// (updates the object on which the method is called)
func (p *person) updateNamePBR2(newName string) {
	p.firstName = newName
}
