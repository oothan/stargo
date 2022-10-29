package main

import "fmt"

func main() {
	square := func(i int) int {
		return i * i
	}

	cubes := func(i int) int {
		return i * i * i
	}

	fmt.Printf("%v \n", calc(square, 9))
	fmt.Printf("%v \n\n", calc(cubes, 5))

	fmt.Printf("%v \n", calc(func(i int) int {
		return i * 2
	}, 3))
}

func calc(f func(int) int, x int) int {
	return f(x)
}
