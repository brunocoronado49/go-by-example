package main

import "fmt"

func main() {
	// Here we create an arra a that will hold
	// exactly 5 ints. The type of elements and length are
	// both part of the array's type. By default an array is
	// zero-valued, which for ints means 0s
	var a [5]int
	fmt.Println("emp:", a)

	// We can set a value at an index using the
	// array[index] = value syntax, and get a value
	// with array[index]
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The builtin let returns a length of an array
	fmt.Println("len:", len(a))

	//
}
