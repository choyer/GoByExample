// `maps` are Go's built-in [associative data type](https://en.wikipedia.org/wiki/Associative_array)
// Some other languages refer to them as `hashes` or `dicts`.
package main

import (
	"fmt"
	"maps"
)

func main() {
	// Create an empty map using the built-in `make`:
	// `make(map[key-type]val-type)`
	m := make(map[string]int)

	// Set key/value pairs using typical `name[key] = val`
	// syntax.
	m["k1"] = 7
	m["k2"] = 13

	// Print a map with `fmt.Println` will show ALL of
	// its key/value pairs.
	fmt.Println("map:", m)

	// Get a value for a key with `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Non-existant key returns the type's zero value
	// https://go.dev/ref/spec#The_zero_value
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// The built-in `len` returns the number of key/value
	// pairs when called on map.
	fmt.Println("len:", len(m))

	// The built-in `delete` removes key/value pairs from
	// a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// To remove *all* key/value pairs from a map, use
	// the `clear` built-in.
	clear(m)
	fmt.Println("map:", m)

	// The optional 2nd return value when getting a value
	// from a map indicates if the key was present in the
	// map. This can be used to disambiguate between
	// mistting keys and keys with zero values such as
	// `0` or `""`. In this case the 1st return value
	// was ignored with the _blank identifier_ `_`
	// as we didn't care about the value of `m["k2"]`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Declare and initialize a new map in a single line.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// The `maps` package contains a number of useful
	// utility functions for maps.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	} else {
		fmt.Println("n != n2")
	}
}
