package main

import "fmt"

//-----------PERSON-----------------
type person struct {
	Name     string
	NickName string
}

func (p person) Talk() {
	fmt.Println("Hello, I'm ", p.Name)
	fmt.Println("My nickname is ", p.NickName)
}

//-----------PERSON-----------------

//-----------CAT------------------
type cat struct{}

func (cat) Talk() { //เหมือนเป็นการระบุว่า นี้เป็น Method ของ cat
	fmt.Println("Nyaa nyaa")
}

//-----------CAT------------------

//-----------DOG------------------
type dog struct{}

func (*dog) Talk() {
	fmt.Println("Wan Wan")
}

//-----------DOG------------------

type talkable interface {
	Talk()
}

func talkWith(t talkable) { //talkWith implement talkable
	t.Talk()
}

func main() {
	p := person{"SubAlgo", "Sual"} //Create p
	talkWith(p)                    //call func(method) talkWith and send paran (p)
	talkWith(cat{})                //call func(method) talkWith and send paran (cat{})
	talkWith(&dog{})               //call func(method) talkWith and send paran (&dog{})
	d := dog{}
	talkWith(&d)
}
