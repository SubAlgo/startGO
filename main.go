package main

import "fmt"

type person struct {
	Name     string
	NickName string
}

func (p person) mutate() {
	p.Name = "Hacker"
	fmt.Println("inside mutate: ", p)
}

func (p *person) mutate2() {
	p.Name = "Hacker"
	fmt.Println("inside mutate2: ", p)
}

func main() {
	var p1 person
	p1.Name = "SubAlgo"
	p1.NickName = "sual"
	fmt.Println(p1)
	p1.mutate()
	fmt.Println("p1 after call method mutate : ", p1)

	fmt.Println("")

	p1.mutate2()
	fmt.Println("p1 after call method mutate2 : ", p1)
}
