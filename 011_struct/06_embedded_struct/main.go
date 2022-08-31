package main

import "fmt"

type generalInfo struct {
	country, hairColor string
}

type player struct {
	name, sport string
	age         int
	info        generalInfo
}

type p struct {
	name, sport string
	age         int
	generalInfo
}

func main() {
	var player1 player
	player1.name = "Ma Ma"
	player1.age = 57
	player1.sport = "Gaming"
	player1.info.country = "Myanmar"
	player1.info.hairColor = "brown"

	fmt.Printf("player1 : %+v \n", player1)

	var player2 p
	player2.name = "Ma Ma"
	player2.age = 57
	player2.sport = "Gaming"
	player2.country = "Myanmar"
	player2.hairColor = "brown"
	fmt.Printf("player2 : %+v \n", player2)

}
