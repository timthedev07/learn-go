package main

import (
	"bufio"
	"fmt"
	"os"
)

// A function that gets the user input from the command line.
func input(str string) string {
	// get input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(str)

	// if condition
	if scanner.Scan() {
		inputStr := scanner.Text()
		return inputStr
	}
	// error handling
	fmt.Println("Failed to get input.")
	return ""
}
