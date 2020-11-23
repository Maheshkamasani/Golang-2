package main

import "fmt"

func add(x int, y int, a int, b int) {
	total := 0
	total = x + y + a + b
	fmt.Println(total)
}

func main() {
	
	add(20, 30 ,40, 50)
}