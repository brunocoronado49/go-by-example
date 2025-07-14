package main

import "fmt"

func main() {
	edades := map[string]int{
		"Bruce":  27,
		"Melina": 2,
		"Seven":  4,
	}

	bruce := edades["Bruce"]
	melina := edades["Melina"]
	seven := edades["Seven"]

	fmt.Println("Bruce:", bruce)
	fmt.Println("Melina:", melina)
	fmt.Println("Seven:", seven)
}
