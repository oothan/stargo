package main

import "fmt"

type rectangle struct {
	w, l int
}

func (c *rectangle) area() int {
	return c.w * c.l
}

func (c *rectangle) perim() int {
	return 2 * (c.l + c.w)
}

type square struct {
	s int
}

func (c *square) area() int {
	return c.s * c.s
}

func (c *square) perim() int {
	return 4 * c.s
}

type shape interface {
	area() int
	perim() int
}

func info(s shape) {
	fmt.Printf("area()=%d perim()=%d \n", s.area(), s.perim())
}

func totalArea(shapes ...shape) int {
	var totalArea int
	for _, s := range shapes {
		totalArea += s.area()
	}
	return totalArea
}
