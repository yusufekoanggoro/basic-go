package main

import "fmt"

func main() {
	/*
		menggunakan var, tanpa tipe data, menggunakan perantara "=" atau ":="
		variable dideklarasikan dengan menggunakan metode type inference
		type inference yaitu metode deklarasi variable yang tipe data-nya diketahui secara otomatis dari data/nilai variable
	*/
	var firstName = "Yusuf"

	lastName := "Eko"
	lastName += " " + "Anggoro"

	fmt.Printf("Hi! My Name %s %s", firstName, lastName)
}
