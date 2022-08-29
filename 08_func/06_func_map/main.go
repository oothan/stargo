package main

import "fmt"

var players = make(map[string]map[string]bool)

func main() {
	addPlayer("Leo", "Messi")
	addPlayer("Roger", "Federer")
	addPlayer("Michael", "Jordan")

	fmt.Println()
	fmt.Println(hasPlayer("Leo", "Messi"))
	fmt.Println(hasPlayer("Roger", "Federer"))
	fmt.Println(hasPlayer("Michael", "Jordan"))
	fmt.Println(hasPlayer("Michael", "Federer"))

	fmt.Println()
	fmt.Println(players)

	fmt.Println()
	for k1, v1 := range players {
		for k2, v2 := range v1 {
			fmt.Println(k1, k2, v2)
		}
	}
}

func addPlayer(fName, lName string) {
	n := players[fName]
	if n == nil {
		n = make(map[string]bool)
		players[fName] = n
	}
	n[lName] = true
}

func hasPlayer(fName, lName string) bool {
	return players[fName][lName]
}
