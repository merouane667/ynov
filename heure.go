package main

import "fmt"

func main() {

	var h, m int
	fmt.Println("entyrez h")
	fmt.Scanln(&h)

	fmt.Println("entyrez m")
	fmt.Scanln(&m)

	if m == 59 {
		m = 0
		h++
	} else if m == 60 {
		m = 1
		h++
	} else {
		m++
	}

	println(h, "heures", m, "minutes")
}
