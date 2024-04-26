package main

import (
	"fmt"
)

// veriadic function. takes any number of trailing arguments
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("total: ", total)
}

func Veriadic() {
	sum(2, 3)
	sum(4, 5)

	nums := []int{2, 3, 4, 5, 6, 7}
	sum(nums...)
}
