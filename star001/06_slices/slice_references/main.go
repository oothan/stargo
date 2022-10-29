package main

import "fmt"

func main() {
	var a []int

	if a == nil {
		fmt.Printf("len(a) = %d\n", len(a))
	}

	b := []int{2, 4, 6, 8, 10}
	if b != nil {
		fmt.Printf("len(b) = %d\n", len(b))
	}

	var x []int
	xLen := len(x)

	fmt.Println(x, len(x), cap(x))

	aa := []int{1000, 2000, 3000}
	if xLen <= len(aa) {
		xLen := len(aa)
		xCap := len(aa) + 1

		x = make([]int, xLen, xCap)
		copy(x, aa)
	}

	fmt.Println(aa, len(aa), cap(aa))
	fmt.Println(x, len(x), cap(x))

}
