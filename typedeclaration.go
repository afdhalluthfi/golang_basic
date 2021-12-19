package main

import "fmt"

func main() {
	type NoKtp string
	type Menikah bool

	var Noktp NoKtp = "34090909090"
	var stsMenikah Menikah = true

	fmt.Println(Noktp)
	fmt.Println(stsMenikah)
}
