package main

import "fmt"

func main() {
	// Underscore (_) adalah reserved variable yang digunakan untuk menampung nilai yang tidak dipakai.
	_ = "oops"
	_ = "salah!"
	name, _ := "Yusuf", "Eko"

	fmt.Printf("Hi! My Name %s", name)
}
