package main

import (
	"fmt"
)

var combinations = make(map[string]Element)
var elements = make(map[string]Element)

type Element struct {
	Id   string
	Name string
}

// Combine tries to combine two elements and returns a match if possible
func (e *Element) Combine(other *Element) (combined Element) {
	combined = combinations[e.Id+"-"+other.Id]
	return
}

func (e *Element) String() string {
	return e.Name
}

// CreateElements will create the test elements
func CreateElements() {
	elem := []Element{
		{"fire", "Fire"},
		{"ice", "Ice"},
		{"steam", "Steam"},
	}

	for _, item := range elem {
		elements[item.Id] = item
	}

	// Make a steam element
	combinations["fire-ice"] = elem[2]
}

func main() {
	CreateElements()

	fmt.Println("Hello.")
	fmt.Printf("Fire is %s\n", elements["fire"])
	fire := elements["fire"]
	ice := elements["ice"]
	fmt.Printf("Combining fire and ice gives: %s\n", fire.Combine(&ice))
}
