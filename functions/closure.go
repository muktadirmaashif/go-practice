package main

import (
	"fmt"
	"strings"
)

func createCustomFilters(filters []string) func(string) bool {
	//
	return func(word string) bool {
		for _, filter := range filters {
			if strings.Contains(word, filter) {
				return true
			}
		}
		return false
	}
}

func Closure() {
	// list of filters
	filters := []string{"bad", "ugly", "sad"}

	// filtering function based on the list
	filterWords := createCustomFilters(filters)

	// checking
	terrible := filterWords("ugly")
	happy := filterWords("happy")
	fmt.Println(terrible, happy)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func seqClosure() {
	nextInt := intSeq()
	// invoking intSeq function each time calling nextInt
	// ++ the variable of outer function using the anonymous function
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// state preserving - another use case
	newInt := intSeq()
	fmt.Printf("New integer value: %d, previous integer value: %d \n", newInt(), nextInt())

}
