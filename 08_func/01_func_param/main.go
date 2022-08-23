package main

import "fmt"

func main() {
	fmt.Println(add(20, 10))

	x1, y1 := swap("Maung", "Hla")
	fmt.Println(x1, y1)

	x2, y2 := swapOne("One", "Two")
	fmt.Println(x2, y2)
}

func add(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func swapOne(x, y string) (r1, r2 string) {
	r1 = y
	r2 = x
	return r1, r2
}
