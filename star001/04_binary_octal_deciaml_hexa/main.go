package main

import "fmt"

func main() {
	fmt.Printf("%b %o %d %x \n", byte('A'), byte('A'), byte('A'), byte('A'))
	fmt.Printf("%b %[1]o %[1]d %[1]x \n", byte('A'))
	fmt.Printf("%b %o %d %x \n", 'A', 'A', 'A', 'A')
}
