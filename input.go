package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// A function that gets the user input from the command line.
func input(str string) (string, error) {
	// get input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(str)

	// if condition
	if scanner.Scan() {
		inputStr := scanner.Text()
		return inputStr, nil
	}
	// error handling
	return "", errors.New("Failed to scan.")
}
