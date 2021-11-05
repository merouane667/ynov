package main

import "fmt"

func main() {
	var nombre, plusGrand, position int
	for i := 0; i < 10; i++ {
		fmt.Println("entrez un nombre ")
		fmt.Scanln(&nombre)

		if plusGrand < nombre {
			plusGrand = nombre
			position = i + 1
		}

	}

	fmt.Println("le plus grand nombre est", plusGrand, "en position", position)
}
