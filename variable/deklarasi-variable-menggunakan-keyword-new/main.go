package main

import "fmt"

type User struct {
	Name string
}

func main() {
	name := new(string) // variable name menampung data bertipe pointer string

	fmt.Println(name) // nilainya alamat memori 0xc0000280a0
	// untuk menampilkan nilainya di-dereference, caranya dengan menuliskan asterisk (*) sebelum nama variable
	fmt.Println(*name) // ""

	// untuk memberi nilainya dengan cara dereference
	*name = "Yusuf Eko Anggoro"
	fmt.Println(*name)

	// dengan struct
	user := new(User)

	(*user).Name = "Yusuf Eko Anggoro Manual" // dengan pointer dereferencing manual
	fmt.Printf("Hi! My Name %s %s", user.Name, "\n")

	user.Name = "Yusuf Eko Anggoro Auto" // dengan pointer dereferencing otomatis untuk field dan method
	fmt.Printf("Hi! My Name %s", user.Name)
}
