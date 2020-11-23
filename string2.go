package main

import "strconv"

func main() {
	strVar := "100"
	intVar, _ := strconv.Atoi(strVar)

	strVar1 := "-78967"
	intVar1, _ := strconv.ParseInt(strVar1, 10, 32)

	strVar2 := "1234567898760"
	intVar2, _ := strconv.ParseInt(strVar2, 10, 64)
}
