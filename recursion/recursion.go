// Go supports [recursive functions](https://en.wikipedia.org/wiki/Recursion_(computer_science))
// Here is a classic example.
package main

import "fmt"

// fact is a function that finds the factorial of the given
// number by calling itself until it reaches the
// base case of `fact(0)`
func fact(n int) int {
	if n == 0 {
		return 1
	}
	r := n * fact(n-1)
	fmt.Println(n, "! =", r)
	return r
}

func main() {
	fmt.Println(fact(7))

	// Anonymous functions can also be recursive, but requires
	// explicitly declaring a variable with `var` to store the
	// function before it's defined.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// since `fib` was previously declared in `main`, Go
		// knows which function to call with `fib` here.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
