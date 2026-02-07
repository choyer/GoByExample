// `slices` are an important data type in Go, providing
// a more powerful interface to sequences than arrays.
package main

import (
	"fmt"
	"slices"
)

func main() {
	// Unlike arrays, slices are typed only by the
	// elements they contain (NOT the number of elements).
	// An uninitialized slice equals nil and has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// To create a slice with non-zero length, use the
	// builtin `make`. Here we make a slice of
	// `string`s of length `3` (initially zero-valued).
	// By default a new slice's capacity is equal to its
	// length; if we know the slice is going to grow
	// ahead of time, it's possible to pass a capacity
	// explicity as an addional paramter to `make`.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// We can set and get like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	// Beyond these basic operations, slices support
	// several more that make them richer than arrays.
	// The builtin `append`, returns a slice containing
	// one or more values.
	// NOTE: you must accept a return value from `append`
	// as it may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be `copy`'d. This example creates
	// an empty slice `c` of the the same length as `s`
	// and copy INTO `c` from `s`.
	// NOTE: `copy` returns the number of elements copied,
	// which will be the min of len(src) && len(dst).
	c := make([]string, len(s))
	r := copy(c, s)
	fmt.Println("cpy:", c)
	fmt.Println("ret:", r)

	// Slices support a "slice" operator with the syntax
	// `slice[low:high]`. ex. this gets a slice of the
	// elements `s[2]`,l `s[3]`, `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l)

	// This slices up from (and including) `s[2]`.
	l = s[2:]
	fmt.Println("sl3:", l)

	// Declare and intitialize a variable for slice
	// in a single line.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// The `slices` package contains a number of useful
	// utility functions for slices.
	t2 := []string{"z", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	} else {
		fmt.Println("t != t2")
	}

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can vary,
	// unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// BONUS! Syntax to create a slice from given array.
	a := [3]string{"Huey", "Dewey", "Louie"}
	x := a[:]
	fmt.Printf("`a` is type %T with length %v\n", a, len(a))
	fmt.Printf("`x` is type %T with length %v and capacity %v. Array Addr is %p\n", x, len(x), cap(x), x)
}
