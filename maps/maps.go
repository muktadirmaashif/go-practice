package main

import (
	"fmt"
	"maps"
)

func main() {
	// make(map[key-type]value-type)
	age := make(map[string]int)

	// name[key]  = value
	age["muk"] = 25
	age["ove"] = 27

	fmt.Println("map: ", age)

	fmt.Println("length of age: ", len(age))

	// delete(age, "ove")
	// clear(age)

	_, prs := age["muk"]
	fmt.Println("prs", prs)

	// Equal
	n1 := map[string]int{"a": 1, "b": 2}
	n2 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(n1, n2) {
		fmt.Println("Equal")
	}
}
