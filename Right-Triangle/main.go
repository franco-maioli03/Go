package main

import (
	"fmt"
	"math"
)

func main() {

	var side1, side2 float64

	fmt.Print("Enter the measurement of the first side: ")
	fmt.Scanln(&side1)
	fmt.Print("Enter the measurement of the second side: ")
	fmt.Scanln(&side2)

	hypotenuse := math.Sqrt(math.Pow(side1, 2) + math.Pow(side2, 2))

	perimeter := side1 + side2 + hypotenuse
	area := (side1 * side2) / 2
	fmt.Printf("\nArea = %.2f", area)
	fmt.Printf("\nPerimeter = %.2f ", perimeter)

}
