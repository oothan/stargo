package main

import (
	"fmt"
	"go-learn/star001/011_struct/08_struct_exported/types"
	"strings"
)

func changeAthleteName1(p athletes.Player) {
	p.Name = "bingo"
}

func changeAthleteName2(p *athletes.Player) {
	p.Name = "bingo"
	p.Country = strings.ToUpper(p.Country)
}

func main() {
	player1 := athletes.Player{Name: "Shwe Mya", Sport: "MMA", Age: 443, Info: athletes.Info{Country: "Brazil", HairCode: "Black"}}
	fmt.Println("(1) player1 : ", player1)

	changeAthleteName1(player1)
	fmt.Println("(2) player2 : ", player1)

	changeAthleteName2(&player1)
	fmt.Println("(3) player3 : ", player1)

	fmt.Println("(4) player4 : ", *player1.ToLowercase())
}
