package main

import "fmt"

func main() {
	player1 := struct {
		name, sport string
		age         int
	}{"Hla Lay", "Gaming", 34}

	fmt.Println("player1 : ", player1)

	player2 := struct {
		name, sport string
		age         int
	}{
		age:   45,
		name:  "Shwe Ba",
		sport: "Drawing",
	}
	fmt.Println("player2 : ", player2)
}
