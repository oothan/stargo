package main

import "fmt"

func main() {
	fmt.Println("=>", factorial(3))
	fmt.Println("=>", factorial(5))
	fmt.Println("=>", factorial(7))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	fmt.Print(num, " ")
	return num * factorial(num-1)
}
