// `Switch` statements express conditionals across many
// branches.
package main

import (
	"fmt"
	"time"
)

func main() {
	// basic switch.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Use commas to seperate multiple expressions
	// in the same `case` statement. Notice the use
	// of the optional `default` case.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend! Yahoo!")
	default:
		fmt.Println("It's a weekday.")
	}

	// `switch` without an expression is an alternate way
	// to express if/else logic. This shows that `case`
	// can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type `switch` compares types instead of values. You
	// can use this to discover the type of an interface
	// value. This example shows that the variable `t`
	// will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Type unkown %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
