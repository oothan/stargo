package main

import "fmt"

func main() {
	//nums := new([3]int)[0:2]
	nums := make([]int, 2, 3)

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	nums[1] = 34
	fmt.Printf("%v\n\n", nums)

	s1 := []int{1, 3, 5}
	s2 := make([]int, 2)

	copy(s2, s1)
	s3 := append(s2, 4, 6)

	fmt.Println(s1, s2, s3)
	fmt.Println(len(s1), len(s2), len(s3))
	fmt.Println(cap(s1), cap(s2), cap(s3))

	fmt.Println()
	for i := range s3 {
		fmt.Print(s3[i], " ")
	}

	fmt.Println()
	for _, el := range s3 {
		fmt.Printf("%d ", el)
	}
	fmt.Println()

	ss1 := []int{2, 3, 4}
	slices := [][]int{ss1, {5, 6}, {7, 8, 9, 10}}
	var bigSlices []int

	for r := range slices {
		bigSlices = append(bigSlices, slices[r]...)
	}

	fmt.Printf("slices = %v %T\n", slices, slices)
	fmt.Printf("bigSlices = %v %T\n", bigSlices, bigSlices)

	x := append(slices[0][1:], slices[2][:1]...)
	fmt.Printf("x = %v %T\n", x, x)
}
