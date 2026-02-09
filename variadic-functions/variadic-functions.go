// [Variadic functions](https://en.wikipedia.org/wiki/Variadic_function)
// can be called with any number of trailing arguments.
// E.g. `fmt.Println` is a commonly used variadic function.
package main

import "fmt"

// sum takes an arbitrary number of `int` arguments and
// prints the total sum.
func sum(nums ...int) {
	fmt.Print("The sum of ", nums)
	total := 0
	// Within the function, the type of `nums` is
	// equivalent to `[]int`. We can call `len(nums)`,
	// iterate over it with `range`, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(" is", total)
}

func main() {
	// Variadic functions can be called in the typical way
	// with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// `func(slice...)` like so:
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
