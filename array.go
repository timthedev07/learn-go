package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Iterable interface{}

func joinArray(arr []string) string {
	str := ""
	// iterate over the array
	for i, n := 0, len(arr); i < n; i++ {
		str = str + arr[i]
	}
	return str
}

func linSearch(arr []int, target int) int {
	for i, item := range arr {
		if item == target {
			return i
		}
	}
	return -1
}

func binSearch(arr []int, target int) int {
	low, high := 0, len(arr)

	for low < high {
		mid := low + ((high - low) / 2)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func doDecentStuffWithArray() {
	min, max, length := 1, 10000000000000, 500000
	rand.Seed(time.Now().UnixNano())
	someArr := []int{}
	for i := 0; i < length; i++ {
		randNumber := rand.Intn(max-min+1) + min
		// append item to slice
		someArr = append(someArr, randNumber)
	}
	// sort an array of integers
	sort.Ints(someArr)

	target := rand.Intn(max-min+1) + min
	fmt.Println("Started searching with linear search.")
	linearSearchResult := linSearch(someArr, target)
	fmt.Println("Started searching with binary search.")
	binarySearchResult := binSearch(someArr, target)
	fmt.Printf("Target %d found at index %d(linear)\n", target, linearSearchResult)
	fmt.Printf("Target %d found at index %d(binary)\n", target, binarySearchResult)

}

func doArrayStuff() {
	// initialize empty integer array of size 5
	var emptyIntArr [5]int
	fmt.Println(emptyIntArr)

	// define string array with some elements
	someStrArray := []string{"Hello ", "world ", "from ", "Go."}
	// gets the length of array just like python
	fmt.Println(len(someStrArray))
	fmt.Println(joinArray(someStrArray))

	// slicing array
	s1 := someStrArray[0:2]
	fmt.Println(s1, "has a capacity of", cap(s1), "and a length of ", len(s1))
	// slicing and retrieving the items that are previously sliced using the cap(capacity) function
	s2 := s1[:cap(s1)]
	fmt.Println(s2, "has a capacity of", cap(s2), "and a length of ", len(s2))
	// iterate using the range keyword
	for i, item := range s2 {
		fmt.Printf("s1[%d] == %s\n", i, item)
	}
}
