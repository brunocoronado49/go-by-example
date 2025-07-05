package main

import "fmt"

func main() {

	// The most basic type, with a simple condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Another way of accomplishing the basic "do this N times"
	// iterations is range over an integer
	for i := range 3 {
		fmt.Println("range", i)
	}

	// Without a condition will loop repeatedly until
	// you break out of the loop or return from the
	// enclosing function
	for {
		fmt.Println("loop")
		break
	}

	// You can also continue to the next interation of the loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
