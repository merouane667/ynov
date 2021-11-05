package main

import (
	"fmt"
)

func main() {
	var min, max int

	fmt.Println("entrez la borne basse")
	fmt.Scanln(&min)

	fmt.Println("entrez la borne haute")
	fmt.Scanln(&max)
	
	fmt.Println("///////")

	for i := min + 2; i < max; i += 2 {
		if min%2 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println(i - 1)
		}
	}
}
