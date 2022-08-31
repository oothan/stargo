package main

import "fmt"

func main() {
	type myType float32
	var total myType

	total = 44
	fmt.Printf("%.2f %T \n", total, total)

	var total1 float64
	total1 = float64(total)
	fmt.Printf("%.2f %T \n", total1, total1)
}
