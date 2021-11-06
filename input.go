package main

import (
	"bufio"
	"fmt"
	"os"
)

func input(str string) string {
	// get input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(str)
	scanner.Scan()
	inputStr := scanner.Text()

	return inputStr
}
