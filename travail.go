package main

import "fmt"

func main() {
	var NHT, PUH, HS10, HS25, HS50, MHS float64

	fmt.Println("entrez le nombre d'heure de travail")
	fmt.Scanln(&NHT)

	fmt.Println("entrez le prix unitaire d'une heure")
	fmt.Scanln(&PUH)

	if NHT <= 40 {

		HS10 = 0

	} else if NHT > 40 {

		if NHT-40 >= 5 {

			HS10 = 5

		} else {

			HS10 = NHT - 40

		}

	}

	if NHT > 45 {

		if NHT-45 >= 5 {

			HS25 = 5

		} else {

			HS25 = NHT - 40 - HS10

		}

	}

	if NHT > 50 {

		HS50 = NHT - 40 - HS25 - HS10
	}

	MHS = (HS10 * PUH * 10 / 100) + (HS25 * PUH * 25 / 100) + (HS50 * PUH * 50 / 100)

	fmt.Println(MHS)

}
