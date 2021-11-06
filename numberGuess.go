package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func playGuessNumberGame() {
	min := 1
	max := 64

	rand.Seed(time.Now().UnixNano())
	randNumber := uint64(rand.Intn(max-min+1) + min)

	// equivalent of `while true {}`
	for {
		userResponse := input(fmt.Sprintf("Guess a number between %d and %d: ", min, max))
		guess, err := strconv.ParseUint(userResponse, 10, 64)
		if err != nil {
			fmt.Println("Failed to extract an integer from your input, try again.")
			continue
		}
		if guess == randNumber {
			fmt.Println("You got it!")
			break
		} else if guess < randNumber {
			fmt.Println("Guess a greater number")
		} else {
			fmt.Println("Guess a smaller number")
		}
		fmt.Println()
	}
}

func playGuessNumberGames() {
	play := true
	for play {
		playGuessNumberGame()
		userResponse := strings.ToLower(input("\n\nPlay again? [Y/n] "))
		if !(userResponse == "y" || userResponse == "yes") {
			fmt.Println("Have a nice day!")
			break
		}
		fmt.Printf("\n\n\n")
	}
}
