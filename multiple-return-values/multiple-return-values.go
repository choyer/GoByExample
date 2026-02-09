// Go has built-in support for multiple return values.
// This feature is used often in idiomatic Go, eg.
// to return both result and error values from a function.
package main

import "fmt"

// vals function returns 2 `int`'s.
func vals() (int, int) {
	return 3, 7
}

func main() {
	// Use 2 different return values from the call
	// with multiple assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// if you only want a subset of the returned values,
	// use the blank identifier `_`.
	_, c := vals()
	fmt.Println(c)
}
