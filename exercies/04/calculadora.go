package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var op string

	fmt.Println("Ingresa un numero:")
	fmt.Scanln(&num1)

	fmt.Println("Ingresa otro numero:")
	fmt.Scanln(&num2)

	fmt.Println("Que operacion quieres hacer (+, -, *, /)?")
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Println("El resultado es:", num1+num2)
	case "-":
		fmt.Println("El resultado es:", num1-num2)
	case "*":
		fmt.Println("El resultado es:", num1*num2)
	case "/":
		fmt.Println("El resultado es:", num1/num2)
	default:
		fmt.Println("Operacion no valida")
	}
}
