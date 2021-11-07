package main

import "fmt"

func someFunction(text string) func() {
	return func() {
		// closure
		fmt.Println(text)
	}
}

func closureStuff() {
	someFunction("Hi")()
}
