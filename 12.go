package main

import "fmt"

func main() {
	var plusGrand, plusPetit, nombreListe, tmp, somme, moyenne, nombre_EP int
	var pourcentage_EP float64
	var nombre int = 1 //pour rentrer dans la condition de la boucle

	for nombre != 0 {
		tmp = nombre
		somme = somme + tmp

		nombreListe++

		fmt.Println("entrez un nombre (zero pour arreter)")
		fmt.Scanln(&nombre)

		if nombre > plusGrand {
			plusGrand = nombre
		}
		if nombre < plusPetit+1 {
			plusPetit = nombre
		}
		if nombre > 0 {
			nombre_EP++
		}

	}
	nombreListe-- // nombre initier a 1
	somme--       // nombre initier a 1

	moyenne = somme / nombreListe

	pourcentage_EP = (float64(nombre_EP) / float64(nombreListe)) * 100

	fmt.Println("le plus grand nombre est", plusGrand, "le plus petit est", plusPetit)
	fmt.Println("le nombre des elements de la liste est", nombreListe, "leurs somme est", somme, "la moyenne est ", moyenne)
	fmt.Println("le nombre des elements positifs de la liste est", nombre_EP, "leurs pourcentages est ", pourcentage_EP, "%")
}
