package main

import "fmt"

func main() {
	fmt.Println(sum(2, 2))
	fmt.Println(sub(2, 3))
	fmt.Println(mult(2, 3))
}

func sum(a int, b int) int {
	num1 := a
	num2 := b

	return num1 + num2
}

func sub(a int, b int) int {
	num1 := a
	num2 := b

	return num1 - num2
}

func mult(a int, b int) int {
	num1 := a
	num2 := b

	return num1 * num2
}
