package main

import "fmt"

func main() {
	//Basic Pointer
	a := 10
	fmt.Println("a = ", a)

	ptrA := &a
	fmt.Println("ptrA = ", ptrA)
	fmt.Println("Value of ptrA = ", *ptrA)

	*ptrA = 20

	fmt.Println(a)

}

/*
Result
a =  10
ptrA =  0xc0420361d0
Value of ptrA =  10
20
*/
