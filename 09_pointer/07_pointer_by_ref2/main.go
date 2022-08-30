package main

import "fmt"

func main() {
	x := 1
	fmt.Printf("(1) x=%d &x=%x\n", x, &x)

	f(&x)
	fmt.Printf("(2) x=%d\n", x)

	fmt.Printf("(3) x=%d\n", f(&x))
}

func f(y *int) int {
	*y++
	fmt.Printf("(f) *y=%d y=%x \n", *y, y)

	return *y
}
