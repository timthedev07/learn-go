package main

import "fmt"

type Iterable interface{}

func joinArray(arr []string) string {
	str := ""
	// iterate over the array
	for i, n := 0, len(arr); i < n; i++ {
		str = str + arr[i]
	}
	return str
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
}
