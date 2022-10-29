package main

import "fmt"

func main() {
	x := []int{2, 4, 6, 8, 10}
	p := &x

	fmt.Printf("x=%v *p=%d &x???=%x p???=%x &p=%x \n", x, *p, &x, p, &p)

	p0 := &x[0]
	fmt.Printf("x[0]=%v *p0=%d &x[0]=%x p0=%x &p0=%x \n", x[0], *p0, &x[0], p0, &p0)

	p2 := &x[2]
	fmt.Printf("x[2]=%v *p2=%d &x[2]=%x p2=%x &p2=%x \n", x[2], *p2, &x[2], p2, &p2)

	fmt.Println()
	y := x[2:]
	q := &y
	fmt.Printf("*q=%d &q=%x q=%T \n", *q, &q, q)

	q0 := &y[0]
	fmt.Printf("y[0]=%v *q0=%d &y[0]=%x q0=%x &q0=%x \n", y[0], *q0, &y[0], q0, &q0)

	fmt.Println()
	(*p)[4] = 100
	for i := range x {
		fmt.Println(i, x[i], (*p)[i])
	}

	fmt.Println()
	(*q)[1] = -9
	y[1] = 700
	for i := range y {
		fmt.Println(i, y[i], (*q)[i])
	}

	t := [...]string{"a", "b", "c", "d", "e", "f"}
	w := &t

	u := t[1:3]
	v := &u

	(*w)[1] = "B"
	(*v)[1] = "C"

	fmt.Println()
	for i := range t {
		fmt.Println(i, t[i], (*w)[i])
	}

	fmt.Println()
	for i := range u {
		fmt.Println(i, u[i], (*v)[i])
	}

}
