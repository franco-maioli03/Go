package animal

import (
	"fmt"
)

type Animal interface {
	Sound()
}

// Dog struct and its methods
type Dog struct {
	Name string
}

func (d Dog) Sound() {
	fmt.Println(d.Name + " says woof woof")
}

// Cat struct and its methods
type Cat struct {
	Name string
}

func (c Cat) Sound() {
	fmt.Println(c.Name + " says meow")
}

// Function to print sound
func MakeSound(animal Animal) {
	animal.Sound()
}
