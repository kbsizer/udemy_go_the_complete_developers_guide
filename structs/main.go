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

	alex.updateName("CHANGED")
	alex.logDetails("after updateName")
}

func (p person) logDetails(msg string) {
	fmt.Printf("%s = %+v\n", msg, p)
}

// pass-by-value (we only update our private copy)
func (p person) updateName(newName string) {
	p.firstName = newName
}

// pass-by-referencee
func (p person) updateName(newName string) {
	p.firstName = newName
}
