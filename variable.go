package main

import "fmt"

func main() {
	var name string

	// jika banyak variable
	var (
		firstname string
		lastname  string
		age       int
		address   string
	)
	firstname = "Afdhal"
	lastname = "Luthfi"
	age = 28
	address = "Bantul"
	name = "Afdhal Luthfi"

	fmt.Println(name)
	fmt.Println(name, firstname, lastname, age, address)
}
