package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 4, 5}
	sum := 0
	for _, num := range nums { // by using _, ignoring the first value. which was in this case, index value
		sum += num
	}
	fmt.Println("sum :", sum)
	// range on arrays and slices provide both the index and value. This time we need the index.
	for i, num := range nums {
		fmt.Println("index of", num, ": ", i)
	}

	kvs := map[string]string{"a": "apple", "b": "ball"}
	for k, v := range kvs {
		fmt.Printf("%s : %s \n", k, v)
	}
	// only key
	for k := range kvs {
		fmt.Printf("%s \n", k)
	}
	// only value, using _, to ignore firest return
	for _, v := range kvs {
		fmt.Printf("%s \n", v)
	}
}
