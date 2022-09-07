package main

import "fmt"

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	a()
	//b()
	c()
}
