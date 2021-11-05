package main

import "fmt"

func main() {

	var nombre, total, tmp int

	fmt.Println("entrez un nombre")
	fmt.Scanln(&nombre)

	tmp = nombre
	for i := 0; i < nombre; i++ {
		total = total + tmp
		tmp--

	}

	fmt.Println(total)
}
