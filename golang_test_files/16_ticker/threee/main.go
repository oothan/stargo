package main

import "fmt"

func main() {
	f := 21.4
	fmt.Println(fmt.Sprintf("%v000000", int64(f)))

	ff := 0.00015
	amount := ff * 1000000000000000000
	fmt.Println(amount)

	//f := ""
}
