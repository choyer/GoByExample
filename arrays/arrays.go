// In Go, an `array` is a numbered sequence of elements of a
// specific length. In typical Go code, `[slices](slices)` are
// much more common; arrays are useful in some special scenarios.
package main

import "fmt"

func main() {
	// An arry `a` that will hold exactly 5 `int`s. The
	// type of elements and length are both part of the
	// array's type. By default an array is zero-valued,
	// which for `int`s means `0`S.
	var a [5]int
	fmt.Println("emp:", a)

	// The value at a specific index can be set using the
	// `array[index] = value` syntax, and the value can be
	// retrieved with `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The built-in `len` returns the length of an array.
	fmt.Println("len:", len(a))

	// Syntax to declare and initialize an array in a
	// single line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// The compiler can infere the number of elements
	// in an array using `...`
	c := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("dcl:", c)
	fmt.Println("len:", len(c))

	// Array types are one-dimensional, however you can
	// compose types to build multi-dimensional data
	// structures.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Create and initialize multi-dimensional arrays
	// at once.
	twoD = [2][3]int{
		{11, 12, 13},
		{21, 22, 23},
	}
	fmt.Println("2d: ", twoD)
}
