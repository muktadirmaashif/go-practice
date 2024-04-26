package main

import (
	"fmt"
)

// func funcName(param paramType) returnType { return something}
// for multiple argument of same type, use (a, b int)
func add(a int, b int) {
	fmt.Println(a + b)
}

func main() {
	// add(3, 2)

	// invoking multiple return function mulRetFunc.go
	// mulRes, cmnt := mulRet(4, 5)
	// fmt.Println("4*5 = ", mulRes, "|| status: ", cmnt)
	//
	// mulRes, cmnt = mulRet(5, 0)
	// fmt.Println("5*0 = ", mulRes, "|| status: ", cmnt)

	// invoking veriadic function from veriadic.go
	// Veriadic()
	//
	// invoking anymous function to form closure from closure.go
	// Closure()
	// seqClosure()

	recursion(2)
	fmt.Print("\n")
	recursion(8)
	fibo(6)

}
