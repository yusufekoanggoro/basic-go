package main

import "fmt"

func main() {
	// cara 1
	var one, two, three string
	one, two, three = "one", "two", "three"

	fmt.Printf("%s %s %s\n", one, two, three)

	// Pengisian nilai juga bisa dilakukan bersamaan pada saat deklarasi
	var four, five, six string = "four", "five", "six"
	fmt.Printf("%s %s %s\n", four, five, six)

	// kalau ingin lebih ringkas
	seven, eight, nine := "seven", "eight", "nine"
	fmt.Printf("%s %s %s", seven, eight, nine)
}
