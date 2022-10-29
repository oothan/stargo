package main

import "fmt"

func main() {
	const (
		a = 2
		b
		c
		d = a * 10
		e
		f
	)
	fmt.Println(a, b, c, d, e, f)

	const (
		a2 = iota + 1
		b2
		c2
		d2 = a2 * 10
		e2
		f2 = a2 + 1
		g2 = iota
	)
	fmt.Println(a2, b2, c2, d2, e2, f2, g2)

	const (
		a3 = iota + 4
		b3 = a3 << 1
		c3 = b3 << 1
		d3 = c3 << 1
		e3 = d3 << 1
		f3 = e3 >> 1
	)
	fmt.Println(a3, b3, c3, d3, e3, f3)
}
