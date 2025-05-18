package main

import "fmt"

func main() {
	name := new(string) // variable name menampung data bertipe pointer string

	fmt.Println(name) // nilainya alamat memori 0xc0000280a0
	// untuk menampilkan nilainya di-dereference, caranya dengan menuliskan asterisk (*) sebelum nama variable
	fmt.Println(*name) // ""

	// untuk memberi nilainya dengan cara dereference
	*name = "Yusuf Eko Anggoro"
	fmt.Println(*name)
}
