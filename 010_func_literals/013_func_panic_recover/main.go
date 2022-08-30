package main

import "fmt"

func main() {

	defer func() {
		errMsg := recover()
		fmt.Println(errMsg)
	}()

	var x map[string]int
	x["key"] = 10
	fmt.Println(x)
}
