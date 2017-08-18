package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("Input num1: ")
	fmt.Scanln(&a)
	fmt.Print("Input num2: ")
	fmt.Scanln(&b)

	r := add(a, b)
	fmt.Println(a, " + ", b, " = ", r)
	fmt.Printf("%d + %d = %d", a, b, r)
}

func add(x int, y int) int {
	return x + y
}
