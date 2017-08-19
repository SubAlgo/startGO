package main

import "fmt"

//Basic struct
type person struct {
	Name     string
	NickName string
}

func main() {
	p1 := person{
		Name:     "SubAlgo",
		NickName: "sugo",
	}
	fmt.Println(p1)
	mutatePerson(p1)
	fmt.Println("p1.name after use func mutatePerson : ", p1.Name)

	mutatePerson2(&p1) //call func mutatePerson2 and send pointer p1
	fmt.Println("p1.name after use func mutatePerson2 : ", p1.Name)
}

func mutatePerson(p person) {
	p.Name = "Hacker"
	fmt.Println("p.Name in func mutatePerson : ", p.Name)
}

func mutatePerson2(p *person) {
	p.Name = "Alola"
	fmt.Println("p.Name in func mutatePerson2 : ", p.Name)
}

/*
Output
{SubAlgo sugo}
p.Name in func mutatePerson :  Hacker
p1.name after use func mutatePerson :  SubAlgo
p.Name in func mutatePerson2 :  Alola
p1.name after use func mutatePerson2 :  Alola
*/
