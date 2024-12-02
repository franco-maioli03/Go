package main

import (
	"fmt"

	"rsc.io/quote"
)

const pi float32 = 3.14

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())

	var (
		firstName, lastName string
		age                 int
	)

	firstName = "Franco"
	lastName = "Maioli"
	age = 21

	fmt.Println(firstName, lastName, age)

}
