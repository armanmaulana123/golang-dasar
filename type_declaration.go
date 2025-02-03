package main

import "fmt"

func main() {
	type NoKTP string

	var ktpArman NoKTP = "1111111"

	var contohKTP string = "2222222"
	fmt.Println(ktpArman)
	fmt.Println(NoKTP(contohKTP))
}
