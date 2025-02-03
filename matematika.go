package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d*e

	fmt.Println(c)

	// Augmented Assignment
	a += 5

	var x = a + b - d*e
	fmt.Println(x)

	// Unary Operator

	var i = 10
	i++
	fmt.Println(i)
}
