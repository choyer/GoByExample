// Go supports constants of character, string, boolean,
// and numeric values.

package main

import (
	"fmt"
	"math"
)

// `const` is used to declare a constant value.
const s string = "constant"

func main() {
	fmt.Println(s)

	// A `const` can also appear inside a function body.
	const n = 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	// a numeric constant has no type until it's given
	// one, such as by an explicity conversion.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. Example, here
	// `math.Sin` expects a `float64`.
	fmt.Println(math.Sin(n))
}
