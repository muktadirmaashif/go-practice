package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func recursion(num int) {
	fmt.Printf("factorial of %d: %d.\n", num, fact(num))

	// recursive closure. needs to be declared with a typed var explicitly
	// can't use := here. cause go won't be able to use it as recursive function inside it.
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Printf("fibonacci of %d: %d.\n", num, fib(num))
}

func fibo(num int) {
	curr, prev := 0, 1
	if num <= 1 {
		fmt.Println(num)
	}
	for i := 2; i <= num; i++ {
		curr, prev = curr+prev, curr
		fmt.Println(curr, prev)
	}
	fmt.Println(curr)

}

// debug : n = 5
// 5 > 2, so won't apply if condition
// so, fib(4)+fib(3)
//    fib(4) = fib(3) + fib(2) = fib(2) + fib(1) + fib(1) =
