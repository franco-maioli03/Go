package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}

func play() {

	random := rand.Intn(100)
	var enterNum int
	var attempts int
	const maxAttempts = 10

	for attempts < maxAttempts {
		attempts++
		fmt.Printf("Enter a number (remaining attempts: %d): ", maxAttempts-attempts+1)
		fmt.Scanln(&enterNum)

		if enterNum == random {
			fmt.Println("¡Congratulations, you guessed the number!")
			playAgain()
			return
		} else if enterNum < random {
			fmt.Println("The number is higher")
		} else if enterNum > random {
			fmt.Println("The number is lower")
		}
	}

	fmt.Println("You haven't more attempts. The number was: ", random)
	playAgain()
}

func playAgain() {
	var choise string
	fmt.Print("¿Do you want to play again? (y/n)")
	fmt.Scanln(&choise)

	switch choise {
	case "y":
		play()
	case "n":
		fmt.Println("¡Thanks for playing!")
	default:
		fmt.Println("Wrong choise. Try again")
		playAgain()
	}

}
