package main

import "fmt"

func main() {
	var note float64

	fmt.Println("entrez votre note")
	fmt.Scanln(&note)

	switch {
	case note >= 0 && note < 10:
		fmt.Println("echec")

	case note < 12:
		fmt.Println("passable")

	case note < 14:
		fmt.Println("assez bien")

	case note < 16:
		fmt.Println("bien")

	case note <= 20:
		fmt.Println("tres bien")

	default:
		fmt.Println("erreur ")
	}
}
