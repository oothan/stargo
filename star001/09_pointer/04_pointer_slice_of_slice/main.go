package main

import "fmt"

func main() {
	x := [...]string{"Angela", "Tim", "Nicole", "Henry"}
	p := &x

	y := x[1:3]
	q := &y

	fmt.Printf("*p=%s &p=%x &x=%x &p=%x\n", *p, &p, &x, p)

	p0 := &x[1]
	fmt.Printf("&x[1]=%x p0=%x *p0=%s\n", &x[1], p0, *p0)

	fmt.Printf("*q=%s &q=%x\n", *q, &q)

	(*p)[1] = "TIM"
	(*q)[1] = "NICOLE"

	fmt.Println()
	for i := range x {
		fmt.Println(i, x[i], (*p)[i])
	}

	fmt.Println()
	for i := range y {
		fmt.Println(i, y[i], (*q)[i])
	}

	fmt.Printf("\n\n")

	slices := [][]int{{1, 2, 3}, {3, 4, 5}, {6, 7, 8, 9}, {10}}
	var bigSlice []int

	var g *[][]int
	g = &slices

	for r := range *p {
		bigSlice = append(bigSlice, (*g)[r]...)
		fmt.Printf("* %d %v %d\n", r, (*g)[r], (*g)[r][0])
	}

	fmt.Printf("%v %T\n", *g, *g)
	fmt.Printf("%v %T\n", bigSlice, bigSlice)
}
