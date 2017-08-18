package main

import "fmt"

func main() {
	// array and slice
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:2]
	fmt.Println("Array a", a)
	fmt.Println("Slice b ", b)
	fmt.Println("")
	b[0] = 20

	fmt.Println("Slice b after edit : ", b)
	b = append(b, 20)
	fmt.Println("Array a after append: ", a)
	fmt.Println("Slice b after append: ", b)
}

//outupt
//Array a [1 2 3 4 5 6 7 8 9 10]
//Slice b  [1 2]

//Slice b after edit :  [20 2]
//Array a after append:  [20 2 20 4 5 6 7 8 9 10]
//Slice b after append:  [20 2 20]
