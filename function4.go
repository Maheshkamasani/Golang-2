package main

import "fmt"

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

func main() {
	var a, p int
	a, p = rectangle(10, 60)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)
}