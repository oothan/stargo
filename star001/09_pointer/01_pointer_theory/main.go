package main

import "fmt"

func main() {
	x := 1
	p := &x
	q := &x

	e := &p

	fmt.Printf("x => type=%T value=%d \n", x, x)
	fmt.Printf("x => type=%T value=%x \n\n", &x, &x)

	fmt.Printf("*p => type=%T value=%d \n", *p, *p)
	fmt.Printf("p => type=%T value=%x \n", p, p)
	fmt.Printf("&p => type=%T value=%x \n\n", &p, &p)

	fmt.Printf("*e => type=%T value=%d \n", *e, *e)
	fmt.Printf("e => type=%T value=%x \n", e, e)
	fmt.Printf("&e => type=%T value=%x \n\n", &e, &e)

	fmt.Printf("*q => type=%T value=%d \n", *q, *q)
	fmt.Printf("q => type=%T value=%x \n", q, q)
	fmt.Printf("&q => type=%T value=%x \n\n", &q, &q)
}
