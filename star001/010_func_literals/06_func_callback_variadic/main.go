package main

import "fmt"

func main() {
	add := func(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}

	fmt.Println(add(1, 2, 4))

	nums := []int{2, 4, 6, 8, 10}
	fmt.Println(add(nums...))

	fmt.Printf("%v \n", calc(add, nums))

	mul := func(nums ...int) int {
		total := 1
		for _, num := range nums {
			total *= num
		}
		return total
	}

	fmt.Printf("%v \n", calc(mul, nums))
}

func calc(f func(...int) int, x []int) int {
	return f(x...)
}
