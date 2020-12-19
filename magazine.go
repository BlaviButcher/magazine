package magazine

import "fmt"

// Subscriber hold a name string a rate float64 and an active boolean
type Subscriber struct {
	name   string
	rate   float64
	active bool
}

// PrintInfo will display all info of the subscriber on seperate lines.
func PrintInfo(s *Subscriber) {
	fmt.Println("Name" + s.name)
	fmt.Println(s.rate)
	fmt.Println(s.active)
}
// DefaultSubscriber takes a string name and creates a subscriber struct using that name and default everything else.
func DefaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}
// ApplyDiscount takes a subscriber and discounts by a dollar.
func ApplyDiscount(s *Subscriber) {
	s.rate = 4.99
}
