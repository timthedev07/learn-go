package main

import (
	"fmt"
	"math/big"
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
	fmt.Println(iterativeFib(150))
	memo := make(map[int]*big.Int)
	fmt.Println(recursiveFib(150, &memo))
}
