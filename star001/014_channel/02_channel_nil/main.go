package main

import "fmt"

func main() {
	var c chan int
	//c := make(chan int)

	if c == nil {
		c = make(chan int)
	}
	fmt.Printf("Type of c: %T\n", c)
	close(c)
}
