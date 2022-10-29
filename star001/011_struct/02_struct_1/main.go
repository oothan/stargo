package main

import "fmt"

type player struct {
	name, sport string
	age         int
}

func main() {
	player1 := player{"MaMa", "Dancing", 21}
	fmt.Println("Player1 : ", player1)

	player2 := player{
		name:  "Shwe Ba",
		sport: "football",
	}
	fmt.Println("Player2 : ", player2)

	player3 := new(player)
	player3.age = 90
	fmt.Println("&Player3 : ", player3)
	fmt.Println("*Player3 : ", *player3)
	fmt.Printf("*player3: %v \n*player3: %+v\n", *player3, *player3)
}
