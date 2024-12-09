package main

import (
	// book "Library/Book"
	"Library/animal"
)

func main() {
	// var myBook = book.NewBook("Pedro Paramo", "Juan Rulfo", 128)
	// var myTextBook = book.NewTextBook("Communication", "Jaime Gamarra", 261,
	// 	"Santilla SAC", "High School")
	// myBook.PrintInfo()
	// myTextBook.PrintInfo()

	// book.Print(myBook)
	// book.Print(myTextBook)

	animals := []animal.Animal{
		&animal.Dog{Name: "Max"},
		&animal.Dog{Name: "Tom"},
		&animal.Cat{Name: "Buddy"},
		&animal.Cat{Name: "Luna"},
	}

	for _, animal := range animals {
		animal.Sound()
	}
}
