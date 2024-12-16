package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // Rock beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // Paper beats rock. (rock + 1) % 3 = 0
	SCISSORS = 2 // Scissors beat paper. (paper + 1) % 3 = 0
)

// Structure to provide the result of each round
type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

// Messages for winning rounds
var winMessages = []string{
	"Well done!",
	"Great job!",
	"You should buy a lottery ticket!",
}

// Messages for losing rounds
var loseMessages = []string{
	"Too bad!",
	"Try again!",
	"Today is just not your day.",
}

// Messages for draws
var drawMessages = []string{
	"Great minds think alike.",
	"Oh oh. Try again.",
	"Nobody wins, but you can try again.",
}

// Variables to keep track of scores
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)
	var computerChoice, roundResult string
	var computerChoiceInt int

	// Message depending on the computer's choice
	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "The computer chose ROCK"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "The computer chose PAPER"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "The computer chose SCISSORS"
	}

	// Generate a random number from 0-2 to choose a random message
	messageInt := rand.Intn(3)

	// Declare a variable to hold the message
	var message string

	if playerValue == computerValue {
		roundResult = "It's a draw"
		// Select a message from drawMessages
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "The Player wins!"
		// Select a message from winMessages
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "The Computer wins!"
		// Select a message from loseMessages
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
