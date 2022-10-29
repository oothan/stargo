package main

import "fmt"

func main() {
	fmt.Println(square(2))
	fmt.Println(square(4))
	fmt.Println(square(3))
}

func square(x int) (result int) {
	result = x * x

	defer func() {
		if x == 2 || x == 4 {
			result += x
		}
	}()

	fmt.Println("* ")
	return
}
