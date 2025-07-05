package main

import "fmt"

func main() {
	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// You can multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// Variables declarated without a corresponding
	// initialization are zero-valued
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and
	// initializing a variable
	// This syntax is only available inside functions
	f := "apple"
	fmt.Println(f)
}
