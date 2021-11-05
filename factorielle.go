package main

import "fmt"

func main() {

	var nombre, tmp int
	var factorielle int = 1

	fmt.Println("entrez un nombre")
	fmt.Scanln(&nombre)

	tmp = nombre
	for i := 1; i <= nombre; i++ {
		factorielle = factorielle * tmp
		tmp--

	}

	fmt.Println(factorielle)
}
