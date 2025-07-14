package main

import "fmt"

func main() {
	var nota1, nota2, nota3, nota4, nota5, prom float64
	var notasTotal []float64

	fmt.Println("Ingresa 5 notas (Enter despues de agregar cada una):")
	fmt.Scanln(&nota1)
	notasTotal = append(notasTotal, nota1)
	fmt.Scanln(&nota2)
	notasTotal = append(notasTotal, nota2)
	fmt.Scanln(&nota3)
	notasTotal = append(notasTotal, nota3)
	fmt.Scanln(&nota4)
	notasTotal = append(notasTotal, nota4)
	fmt.Scanln(&nota5)
	notasTotal = append(notasTotal, nota5)

	for i := range notasTotal {
		prom += notasTotal[i]
	}

	prom = prom / float64(len(notasTotal))
	fmt.Println("Calif final:", prom)
}
