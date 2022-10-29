package main

import "fmt"

func main() {
	var nums [3]int

	var sum1 int
	var sum2 int

	for j := range nums {
		sum1 += j
		sum2 += nums[j]
	}

	fmt.Println(sum1, sum2)

	x := [3]float32{2.1, 3.2, 4.7}

	fmt.Println(x)
	var total float32

	for _, val := range x {
		total += val
	}

	fmt.Println(total)

	x1 := []int{10, 15, 20}

	x2 := [...]int{10, 15, 20, 25}
	x3 := [4]int{10, 15, 20, 25}
	var x4 [4]int = [4]int{10, 15, 20, 25}

	//fmt.Println(x1, x1 == x2, x2 == x3, x2 == x4, x3 == x4)
	fmt.Println(x1, x2 == x3, x2 == x4, x3 == x4)

	// custom array index
	type Size int

	const (
		EX Size = iota + 1
		LG
		MD
		SM
	)

	sz := [...]string{EX: "Extra Large", LG: "Large", MD: "Medium", SM: "Small"}

	fmt.Println(MD, sz[MD])
	fmt.Println(sz[2] == "Large", "Large")
	fmt.Println(sz)

	// array index make
	xx := [...]int{10, 10: 20}
	fmt.Println(xx)

	xx1 := [...]int{3: 10, 20}
	fmt.Println(xx1)

	xx2 := [...]int{3: 10, 10: 20}
	fmt.Println(xx2)

	// multi dimension array
	a := [2][3]int{{2, 4, 6}, {4, 6, 8}}
	fmt.Printf("a = %v type=%T, len=%d\n", a, a, len(a))

	var b [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = (i + j + 1) * 2
		}
	}
	fmt.Printf("b = %v type=%T, len=%d\n", b, b, len(b))

	a[0][1] = 5
	a[1][1] = 10
	fmt.Printf("a = %v type=%T, len=%d\n", a, a, len(a))

}
