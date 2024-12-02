package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())

	var firstName, lastName string
	var age int

	firstName = "Franco"
	lastName = "Maioli"
	age = 21

	fmt.Println(firstName, lastName, age)

}
