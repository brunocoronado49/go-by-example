package main

// Range iterates over elements in a variety of built-in data
// structures. Let's see how to use range with some of the data
// structures we've already learned
func main() {
	// Here we use range to sum the numbers in a slice. Array work
	// like this too
	nums := []int{2, 4, 6}
	sum := 0
	for _, num := range nums {
		sum += num
	}
}
