package main

import (
	"fmt"
)

type artist struct {
	name string
	age  int
}

func (a *artist) showName() {
	fmt.Println("Name : ", a.name)
}

func (a *artist) showAge() {
	fmt.Println("Age : ", a.age)
}

type singer struct {
	field string
	artist
}

func (a *singer) showAge() {
	fmt.Println("Fake age : ", a.age+100)
}

func (a *singer) showField() {
	fmt.Println("Field : ", a.field)
}

func main() {
	r1 := artist{
		name: "Go Go",
		age:  22,
	}
	fmt.Println(r1)
	r1.showName()
	r1.showAge()

	s1 := singer{
		field:  "singer",
		artist: r1,
	}

	fmt.Println(s1)
	s1.showField()
	s1.showAge()
	s1.showName()

	s2 := singer{
		artist: artist{
			name: "Doggy",
			age:  102,
		},
		field: "singer",
	}
	fmt.Printf("\n%v\n", s2)
	s2.showField()
	s2.showName()
	s2.showAge()

}
