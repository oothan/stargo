package main

import "fmt"

func main() {
	defer fmt.Printf("#A main \n")
	defer fmt.Printf("#B main \n")
	callPop()
}

func callPop() {
	defer fmt.Printf("#C callPop\n")
	pop(3)
}

func pop(x int) {
	fmt.Printf("pop(%d) \n", 10/x)
	defer fmt.Printf("#%d pop \n", x)
	pop(x - 1)
}
