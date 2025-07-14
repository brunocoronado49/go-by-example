package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var vocals int

	fmt.Println("Ingresa un algo...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cadena := scanner.Text()

	cadena = strings.ToLower(cadena)

	for _, c := range cadena {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vocals++
		}
	}

	fmt.Println("El total de vocales es:", vocals)
}
