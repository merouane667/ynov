package main

import "fmt"

func main() {
	var nhs, mp, mar, vmp, r int
	var p10, p5, p2, p1 int
	var reprendre bool = true
	fmt.Println("entrez le nombre d'heure de stationnement en minutes : ")
	fmt.Scanln(&nhs)

	if nhs >= 30 {

		for {
			vmp = (nhs * 2) / 60

			if reprendre {

				fmt.Println("vous devez payer : ", vmp, "dh")
				fmt.Println("tapez un montant")
				fmt.Scanln(&mp)
			}

			if mp-vmp < 0 {

				fmt.Println("vous devez payer : ", vmp, "dh")
				fmt.Println("tapez un montant")
				fmt.Scanln(&mp)

				reprendre = false

			} else {

				mar = mp - vmp
				fmt.Println("le montant a rendre est : ", mar)

				p10 = mar / 10
				r = mar % 10
				p5 = r / 5
				r = r % 5
				p2 = r / 2
				r = r % 2
				p1 = r

				fmt.Println("sous forme de ", p10, "pieces de 1o dh et", p5, "pieces de 5 dh et", p2, "pieces de 2 dh et", p1, "pieces de 1 dh")

				break
			}
		}

	} else {

		fmt.Println("vous n'avez pas depasse une demie heure , votre stationnement est gratuit  ")

	}

}
