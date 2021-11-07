package main

import (
	"fmt"
	"math"
)

func main() {
	// casting, using the math module, and printing decimals
	x := 5
	calcResult := math.Sin(1/(math.Pow(float64(x)+3.482847, 3))) * (180 / math.Pi)
	fmt.Printf("%g\n", calcResult)
	fmt.Printf("%F\n", calcResult)
	fmt.Printf("%f\n", calcResult)
	fmt.Printf("%.6f\n", calcResult)
	fmt.Printf("%e\n", calcResult)

	// format string
	someStr := fmt.Sprintf("Calculation result: %f", calcResult)
	someOtherStr := fmt.Sprintf("Calculation result to 4dp: %.4f", calcResult)
	fmt.Println(someStr)
	fmt.Println(someOtherStr)
	fmt.Printf("\n\n")

	// play the number guessing game
	// playGuessNumberGames()
	doArrayStuff()
	doDecentStuffWithArray()

	doMapStuff()
}
