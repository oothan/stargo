package main

import "fmt"

func main() {
	x := 1
	p := &x

	fmt.Printf("x=%T &x=%T p=%T *p=%T &p=%T \n", x, &x, p, *p, &p)
	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x \n\n", x, &x, p, *p, &p)

	*p = 2
	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x \n\n", x, &x, p, *p, &p)

	x = 90
	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x \n\n", x, &x, p, *p, &p)

	y := 4
	p = &y
	fmt.Printf("y=%d &y=%x p=%x *p=%d &p=%x \n\n", y, &y, p, *p, &p)
	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x \n\n", x, &x, p, *p, &p)

	q := new(int) // *int
	fmt.Println()
	fmt.Printf("*q=%d q=%x \n", *q, q)

	*q = 10
	fmt.Printf("*q=%d q=%x \n", *q, q)
}
