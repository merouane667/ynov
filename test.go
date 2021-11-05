package main

import "fmt"

func main() {
	var montant, b200, b100, b50, b20, p10, p5, p2, p1, reste int

	fmt.Println("entrez un montant a retirer")
	fmt.Scanln(&montant)

	b200 = montant / 200
	reste = montant % 200
	b100 = reste / 100
	reste %= 100
	b50 = reste / 50
	reste %= 50
	b20 = reste / 20
	reste %= 20
	p10 = reste / 10
	reste %= 10
	p5 = reste / 5
	reste %= 5
	p2 = reste / 2		
	reste %= 2
	p1 = reste

	fmt.Println(b200, "billets de 200dh et", b100, "billets de 100dh et", b50, "billets de 50dh et", b20, "billets de 20dh et ",
		p10, "pieces de 10dh et", p5, "pieces de 5dh et", p2, "pieces de 2dh et", p1, "pieces de 1dh ")

}
