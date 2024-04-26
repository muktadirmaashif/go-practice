package main

import "fmt"

func zeroval(ival int) int {
	ival = 0
	return ival
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("i is ", i)

	// zeroval doesn't change i's value, cause we didn't print its returned var.
	zeroval(i)
	fmt.Println("after zeroval, i is ", i)

	// but zeroptr changes the i in the main, cause called with it's memory address.
	zeroptr(&i)
	fmt.Println("after zeroptr, i is ", i)
	fmt.Println("addr of i is ", &i)

}
