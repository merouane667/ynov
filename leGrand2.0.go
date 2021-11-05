package main

import "fmt"

func main() {
	var nombre, plusGrand, position int
	var i int = 1
	fmt.Println("entrez un nombre (zero pour arreter")
	fmt.Scanln(&nombre)

	for nombre != 0 {
		fmt.Println("entrez un nombre (zero pour arreter)")
		fmt.Scanln(&nombre)

		if plusGrand < nombre {
			plusGrand = nombre
			i++
			position = i
		}

	}

	fmt.Println("le plus grand nombre est", plusGrand, "en position", position)
}
