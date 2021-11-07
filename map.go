package main

import (
	"fmt"
	"strings"
)

// 															 labeling return values
func findDuplicates(data []int) (duplicate []int) {
	seen := make(map[int]uint8)
	duplicate = []int{}
	for _, item := range data {
		times, hasItem := seen[item]
		if hasItem && times == 1 {
			duplicate = append(duplicate, item)
		} else {
			seen[item] = 1
		}
	}
	return
}

func doMapStuff() {
	// define a map
	scores := map[string]int{
		"Tom":   94,
		"Jeff":  87,
		"Chris": 93,
		"Brown": 100,
	}
	// brown cheated, give him a zero
	scores["Jeff"] = 0
	// Louis also did the test add him to the map
	scores["Louis"] = 97
	// actually brown sucks, so delete him
	delete(scores, "Brown")
	// another way to define map
	emptyMap := make(map[string]string)
	fmt.Println(emptyMap)
	fmt.Println(scores)

	duplicates := findDuplicates([]int{1, 2, 7, 4, 2, 8, 5, 4, 9, 7})
	numDuplicates := len(duplicates)
	if numDuplicates > 0 {
		var output strings.Builder
		output.WriteString("Found duplicates [")
		for _, duplicate := range duplicates[:numDuplicates-1] {
			output.WriteString(fmt.Sprintf("%d, ", duplicate))
		}
		output.WriteString(fmt.Sprintf("%d]\n", duplicates[numDuplicates-1]))
		fmt.Println(output.String())
	}
}
