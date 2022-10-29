package main

import "fmt"

func main() {
	//var i int
	//for i = 0; i < 15; i++ {
	//	fmt.Printf("%d ", fibonacci(i))
	//}

	fmt.Printf("\n\n%d", fibonacci(6))
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}
