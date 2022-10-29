package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{56, 9, 0, -3, 8, 45}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)

	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	//sort.Ints(n)
	fmt.Println(n)
}
