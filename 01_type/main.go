package main

import "fmt"

func main() {
	firstName := "Hla Hla"
	firstName = "Maung"
	fmt.Println(firstName)

	var i int
	fmt.Println(" i = ", i)

	j := 2
	j++
	fmt.Printf("%d - %T\n", j, j)

	var k uint8 = 20
	fmt.Printf("%d - %T\n", k, k)

	var k2 uint16 = 30
	fmt.Printf("%d - %T\n", k2, k2)

}
