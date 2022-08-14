package main

import "fmt"

func main() {
	//var k1 uint8 = 20

	f1 := 323.14
	fmt.Printf("%f - %T\n", f1, f1)
	fmt.Printf("%5.3f %.2f\n", f1, 214.437)

	type myString string
	var mystr myString = "Hello "

	fmt.Printf("%v %T", mystr, mystr)
}
