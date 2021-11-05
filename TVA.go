package main

import "fmt"

func main() {
	var P_unitaire, Quantite, TauxRemise, Remise, TotalAvantRemise, TotalHT, MontantTVA, TotalTTC float64
	const TVA float64 = 0.2

	fmt.Println("entrez le prix unitaires")
	fmt.Scanln(&P_unitaire)

	fmt.Println("entrez la quantite")
	fmt.Scanln(&Quantite)

	fmt.Println("entrez le taux de la remise (exp: si 10% tapez 0.1) ")
	fmt.Scanln(&TauxRemise)

	TotalAvantRemise = P_unitaire * Quantite
	Remise = TauxRemise * TotalAvantRemise
	TotalHT = TotalAvantRemise - Remise
	MontantTVA = TotalHT * TVA
	TotalTTC = TotalHT * MontantTVA

	fmt.Println("le total avant remise est : ", TotalAvantRemise)
	fmt.Println("le total hors taxe : ", TotalHT)
	fmt.Println("le montant tva : ", MontantTVA)
	fmt.Println("le total ttc  : ", TotalTTC)

}
