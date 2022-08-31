package main

import "fmt"

type movie struct {
	name  string
	actor string
}

func (m movie) fullInfo() string {
	return m.name + " " + m.actor
}

func main() {
	m1 := movie{name: "Forest Gump", actor: "Tom Hunks"}
	fmt.Println(m1.fullInfo())
	fmt.Printf("%+v \n", m1.fullInfo())
}
