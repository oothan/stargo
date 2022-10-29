package main

import "fmt"

func main() {
	nextPos := getPositiveInt()

	fmt.Println(nextPos())
	fmt.Println(nextPos())
	fmt.Println(nextPos())

	fmt.Println()
	nextPos1 := getPositiveInt()
	fmt.Println(nextPos1())

	fmt.Printf("'%T' '%T' \n", nextPos, nextPos())
	fmt.Printf("%x %x \n", &nextPos, &nextPos1)
}

func getPositiveInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
