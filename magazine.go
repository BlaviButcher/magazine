package magazine

import "fmt"

// Subscriber hold a name string a rate float64 and an active boolean
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

// PrintInfo will display all info of the subscriber on seperate lines.
func PrintInfo(s *Subscriber) {
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Rate: %0.2f\n", s.Rate)
	fmt.Printf("Active: %t\n", s.Active)
}
// DefaultSubscriber takes a string name and creates a subscriber struct using that name and default everything else.
func DefaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}
// ApplyDiscount takes a subscriber and discounts by a dollar.
func ApplyDiscount(s *Subscriber) {
	s.Rate = 4.99
}

// Employee holds a name string value and salary float64 value
type Employee struct {
	Name string
	Salary float64
	// use annonymous address so that we don't need to reference variable. This makes access less tedious
	Address
}

// Address holds the values that build up a home address
type Address struct {
	Street string
	City string
	State string
	PostalCode string
}


