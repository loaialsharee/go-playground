package main

import "fmt"

// Animal interface (defines behavior but not implementation)
type Animal interface {
	MakeSound() // Any type implementing this must define MakeSound()
}

// Dog struct (a type that implements Animal interface)
type Dog struct{}

// MakeSound for Dog (implements the Animal interface)
func (d Dog) MakeSound() {
	fmt.Println("Woof! Woof!")
}

// Cat struct (another type that implements Animal interface)
type Cat struct{}

// MakeSound for Cat (implements the Animal interface)
func (c Cat) MakeSound() {
	fmt.Println("Meow! Meow!")
}

// Speak function (accepts any Animal and calls MakeSound)
func Speak(a Animal) {
	a.MakeSound()
}

func main() {
	dog := Dog{}
	cat := Cat{}

	// Both Dog and Cat implement the Animal interface, so they can be used in Speak
	Speak(dog) // Output: Woof! Woof!
	Speak(cat) // Output: Meow! Meow!
}
