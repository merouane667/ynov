package main

import "fmt"

func main() {
	var HT, S, PS, r float64
	fmt.Println("Donner le prix unitaire d'heure :")
	fmt.Scanln(&PS)
	fmt.Println("Donner le nombre d'heures de travail de l'employé :")
	fmt.Scanln(&HT)
	switch {
	case 0 <= HT && HT <= 40:
		S = 0
	case 41 <= HT && HT <= 45:
		S = (HT - 40) * 0.1 * PS
		r = S
	case 46 <= HT && HT <= 50:
		S = r + (HT-46)*0.25*PS
		r = S
	case 51 <= HT:
		S = r + (HT-51)*0.5*PS
	default:
		fmt.Println("Valeur invalide")

	}
	fmt.Println("Le montant des heures supplémentaires de l'employé est :", S, "DHS")
}
