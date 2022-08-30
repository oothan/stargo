package main

import "fmt"

func main() {
	var p = f(2)
	fmt.Printf("p=%x *p=%d &p=%x\n", p, *p, &p)

	var q = f(90)
	fmt.Printf("q=%x *q=%d &q=%x\n", q, *q, &q)
}

func f(inp int) *int {
	v := inp * 2
	fmt.Printf("&v=%x\n", &v)

	return &v
}
