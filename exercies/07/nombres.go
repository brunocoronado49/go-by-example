package main

import "fmt"

func main() {
	names := []string{"Francisco", "Bruno", "Eve", "Lady Maria", "Melina"}

	for i := range names {
		fmt.Println("Hola", names[i])
	}
}
