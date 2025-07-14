package main

import "fmt"

func main() {
	var num int

	fmt.Println("Ingresa el numero de la tabla:")
	fmt.Scanln(&num)

	for i := 0; i <= 10; i++ {
		result := num * i
		fmt.Printf("%v * %v = %v\n", num, i, result)
	}
}
