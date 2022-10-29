package main

import "fmt"

func main() {
	addCounter, mulCounter := addBy(), mulBy()

	fmt.Print(addCounter(2), " ")
	fmt.Print(addCounter(3), " ")
	fmt.Print(addCounter(-10), "\n")

	fmt.Print(mulCounter(3), " ")
	fmt.Print(mulCounter(4), " ")
	fmt.Print(mulCounter(-2), "\n")
}

func addBy() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func mulBy() func(int) int {
	total := 1
	return func(i int) (ret int) {
		total *= i
		ret = total
		return
	}
}
