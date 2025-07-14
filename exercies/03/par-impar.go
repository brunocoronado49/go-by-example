package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingresa un numero para decir si es par o impar:")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}
}
