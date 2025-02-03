package main

import "fmt"

func main() {
	var name string
	name = "Arman Maulana Saputra"
	fmt.Println(name)

	name = "Maulana"
	fmt.Println(name)

	var age = 22
	fmt.Println(age)

	age = 23
	fmt.Println(age)

	pet := "cat"
	fmt.Println(pet)

	pet = "dog"
	fmt.Println(pet)

	var (
		address = "Jl. Kintamani"
		city    = "Lumajang"
	)

	fmt.Println(address)
	fmt.Println(city)

	const gender string = "male"
	fmt.Println(gender)

	const (
		firstName string = "Arman"
	)

	fmt.Println(firstName)
}
