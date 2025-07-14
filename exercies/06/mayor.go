package main

import "fmt"

func main() {
	var num1, num2, num3 int
	var numMayor int

	fmt.Println("Ingresa 3 numeros, presiona enter despues de agregar cada uno:")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanln(&num3)

	if num1 > num2 {
		numMayor = num1
	} else {
		numMayor = num2
	}

	if numMayor > num3 {
		fmt.Println("El numero mayor es:", numMayor)
	} else {
		fmt.Println("El numero mayor es:", num3)
	}
}
