package main

import (
	"fmt"
)

func main() {
	var nombre1, nombre2, resultat int

	fmt.Println("entrez le premier nombre")
	fmt.Scanln(&nombre1)

	fmt.Println("entrez le deuxieme nombre")
	fmt.Scanln(&nombre2)

	for i := 0; i < nombre2; i++ {
		resultat = resultat + nombre1
	}
	fmt.Println(resultat)

}
