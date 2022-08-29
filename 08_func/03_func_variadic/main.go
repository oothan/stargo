package main

import "fmt"

func main() {
	sum(1, 2, 3, 4, 6)
	sum(45, 89)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum(nums...)
	sum2(nums)

	fmt.Printf("sum=%T sum2=%T \n", sum, sum2)
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Printf("nums=%v total=%d type=%T \n", nums, total, nums)
}

func sum2(nums []int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("nums=%v total=%d type=%T \n", nums, total, nums)
}
