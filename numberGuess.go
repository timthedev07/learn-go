package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func playGuessNumberGame() {
	min := 1
	max := 20

	randNumber := uint64(rand.Intn(max-min+1) + min)

	// equivalent of `while true {}`
	for {
		userResponse := input("Guess the number: ")
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
	}
}

func playGuessNumberGames() {
	play := true
	for play {
		playGuessNumberGame()
		userResponse := strings.ToLower(input("Play again?[Y/n] "))
		if !(userResponse == "y" || userResponse == "yes") {
			fmt.Println("Have a nice day!")
			break
		}
	}
}
