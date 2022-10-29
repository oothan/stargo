package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("x=%d \n", x)

	f(&x)
	fmt.Printf("x=%d \n", x)
}

func f(in *int) {
	fmt.Printf("(f) in=%d\n", *in)

	*in += 90
	fmt.Printf("(f) in=%d \n", *in)
}
