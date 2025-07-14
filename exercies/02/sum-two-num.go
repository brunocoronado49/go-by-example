package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Println("Ingresa un numero:")
	fmt.Scanln(&num1)

	fmt.Println("Ingresa otro numero:")
	fmt.Scanln(&num2)

	fmt.Println("La suma es:", num1+num2)
}
