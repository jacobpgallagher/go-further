package main

import "fmt"

// START OMIT
type Animal struct{ Name string }

func (a *Animal) Eat() {
	fmt.Println("I don't know what to eat. I'll eat dirt.")
}

func (a *Animal) Speak() {
	fmt.Println("...")
}

type Herbivore struct{ *Animal }

func (h *Herbivore) Eat() {
	fmt.Println("Mmm. Plants. Tasty.")
}

type Cow struct{ *Herbivore }

func (c *Cow) Speak() {
	fmt.Printf("Moo! (Also my name is %s)\n", c.Name)
}

// START MAIN OMIT
func main() {
	bob := Animal{"Bob"}
	bob.Eat()
	bob.Speak()

	alice := Herbivore{&Animal{"Alice"}}
	alice.Eat()
	alice.Speak()

	bessy := Cow{&Herbivore{&Animal{"Bessy"}}}
	bessy.Eat()
	bessy.Speak()

	// Probably best to not name things that people will kill and eat.
	bessy.Name = "Cow"
	bessy.Speak()

	bessy.Herbivore.Animal.Eat()
}

//END OMIT
//END MAIN OMIT
