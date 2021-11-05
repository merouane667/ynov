package main

import "fmt"

func main() {

	var j, m, a int
	var finAnnee, finmois, finmois31, finmois30, finmois29, finmois28 bool

	fmt.Println("entrez le jour")
	fmt.Scanln(&j)

	fmt.Println("entrez le mois")
	fmt.Scanln(&m)

	fmt.Println("entrez l'annee")
	fmt.Scanln(&a)

	finAnnee = j == 31 && m == 12

	finmois31 = j == 31 && (m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10)
	finmois30 = j == 30 && (m == 4 || m == 6 || m == 9 || m == 11)
	finmois29 = j == 29 && m == 2
	finmois28 = j == 28 && m == 2 && a%4 != 0

	finmois = finmois31 || finmois30 || finmois29 || finmois28

	switch {

	case finAnnee:
		j, m, a = 1, 1, a+1

	case finmois:
		j, m, a = 1, m+1, a

	default:
		j, m, a = j+1, m, a
	}

	fmt.Println(j, "/", m, "/", a)

}
