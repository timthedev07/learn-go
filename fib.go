package main

import (
	"fmt"
	"math/big"
	"time"
)

func iterativeFib(n int) *big.Int {
	fib := []*big.Int{big.NewInt(1), big.NewInt(1)}

	for i := 2; i < n; i++ {
		fib = append(fib, big.NewInt(0).Add(fib[i-1], fib[i-2]))

	}

	return fib[n-1]
}

// memoization with big integers and pointers
func recursiveFib(n int, memo *(map[int]*big.Int)) *big.Int {
	if n <= 2 {
		return big.NewInt(1)
	}
	memoizedValue, memoized := (*memo)[n]
	if memoized {
		return memoizedValue
	}

	fibResult := big.NewInt(0).Add(recursiveFib(n-1, memo), recursiveFib(n-2, memo))

	(*memo)[n] = fibResult

	return fibResult
}

func doStuffWithFib() {
	n := 300000
	start := time.Now()
	iterativeFib(n)
	iterDuration := time.Since(start)

	memo := make(map[int]*big.Int)
	start = time.Now()
	recursiveFib(n, &memo)
	recurDuration := time.Since(start)

	fmt.Printf("The iterative approach took %d ms.\n", iterDuration.Milliseconds())
	fmt.Printf("The recursive approach took %d ms.\n", recurDuration.Milliseconds())

	// lambda function
	fmt.Println(func() string {
		return "Hello from closure."
	}())
}
