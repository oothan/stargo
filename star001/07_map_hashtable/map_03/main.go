package main

import "fmt"

func main() {
	days := make(map[string]int)

	fmt.Println(days)

	days["Sun"] = 3
	days["Sun"] -= 2
	fmt.Println(days)

	days["Mon"] = 1
	days["Mon"]++
	fmt.Println(days)

	fmt.Println(days["Mon"])
	fmt.Println(days["NoDay"])

	powers := make(map[int]int)
	powers[2] = 4
	powers[3] = 9
	fmt.Println(powers)
	fmt.Println(powers[3])
	fmt.Println(powers[1])

	delete(powers, 1)
	fmt.Println(powers)
	delete(powers, 2)
	fmt.Println(powers)

	powers[3] = 10
	powers[1] = 20
	powers[4] = 90
	fmt.Println(powers)

}
