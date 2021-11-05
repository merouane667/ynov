package main

import (
	"fmt"
)

func main() {
	var n, somme_carres, tmp int

	fmt.Println("entrez un entier naturel")
	fmt.Scanln(&n)

	tmp = n
	for i := 1; i <= n; i++ {
		somme_carres = somme_carres + tmp*tmp
		tmp--

	}

	fmt.Println("la somme des carres naturel est : ", somme_carres)

}
