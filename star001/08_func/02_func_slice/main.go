package main

import "fmt"

func main() {
	scores1 := []float32{20, 90, 67, 89}
	fmt.Printf("average: %.2f\n", avg(scores1))

	scores2 := []float32{34, 56, 90, 34, 23, 45, 78, 90, 89, 67}
	fmt.Printf("average: %.2f\n", avg(scores2))
}

func avg(scores []float32) float32 {
	var total float32
	for _, score := range scores {
		total += score
	}

	return total / float32(len(scores))
}
