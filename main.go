package main

import (
	"errors"
	"fmt"
)

var (
	errAgeTooLow  = errors.New("age too low")
	errAgeTooHigh = errors.New("age too high")
)

func validateAge(age int) error {
	if age < 18 {
		return errAgeTooLow
	} else if age > 60 {
		return fmt.Errorf("age too high")
	} else {
		return nil
	}
}

func main() {

	var age int
	fmt.Print("Input your age: ")
	fmt.Scanln(&age)

	err := validateAge(age)
	if err == errAgeTooLow {
		fmt.Println("CAN NOT ENTER")
		return
	}
	if err != nil {
		fmt.Println("fuuuu")
		return
	}
}
