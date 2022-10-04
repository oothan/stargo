package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi("1" + nchars("0", 5))
	fmt.Println(n)
}

func nchars(b string, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += b
	}
	return s
}
