package main

import "fmt"

func main() {

	var h, m int

	fmt.Println("entrez le nombre d'heures")
	fmt.Scanln(&h)

	fmt.Println("entrez le nombre de minutes")
	fmt.Scanln(&m)

	if m == 59 {
		m = 0
		h = h + 1

	} else {
		m = m + 1
	}

	fmt.Println(h, ":" , m)

}
