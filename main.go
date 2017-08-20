package main

//Basic GoRutines and Chanel
import (
	"fmt"
	"time"
)

var (
	arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
)

func main() {
	fmt.Println("WithGo")
	withGo()
	fmt.Println("WithOutGo")
	withOutGo()
}

//function withGo
func withGo() {
	defer timer()()
	chRes1 := make(chan int) //Create Chanel
	chRes2 := make(chan int) //Create Chanel

	go func() {
		chRes1 <- sum(arr1)
	}()

	go func() {
		chRes2 <- sum(arr2)
	}()

	fmt.Println("Sum arr1: ", <-chRes1)
	fmt.Println("Sum arr2: ", <-chRes2)
}

//Function withOutGo
func withOutGo() {
	defer timer()()
	fmt.Println("Sum arr1: ", sum(arr1))
	fmt.Println("Sum arr2: ", sum(arr2))
}

//Function timmer
func timer() func() {
	t := time.Now()
	return func() {
		diff := time.Now().Sub(t)
		fmt.Println(diff)
	}
}

//Function sum
func sum(a []int) int {
	s := 0
	for _, x := range a {
		s += x
		time.Sleep(time.Millisecond * 200)
	}
	return s
}
