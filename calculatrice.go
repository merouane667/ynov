package main

import "fmt"

func main() {
	var nb1, nb2, resultat float64
	var operation string

	fmt.Println("entrez le premier nombre ")
	fmt.Scanln(&nb1)

	fmt.Println("entrez le deuxieme nombre ")
	fmt.Scanln(&nb2)

	fmt.Println("entrez le type d'operation ( + - * / ) ")
	fmt.Scanln(&operation)

	if operation == "+" {
		resultat = nb1 + nb2
	}
	if operation == "-" {
		resultat = nb1 - nb2
	}
	if operation == "*" {
		resultat = nb1 * nb2
	}
	if operation == "/" {
		resultat = nb1 / nb2
	}

	fmt.Println("le resultat est", resultat)
}
