package main

import "fmt"

func main() {
	printMsg("calling a Function!")

	showMsg := func(msg string) {
		fmt.Println(msg)
	}

	showMsg("Using a function literal!")
	fmt.Printf("%T\n", showMsg)
	fmt.Printf("%T\n", printMsg)

	func(msg string) {
		fmt.Println(msg)
	}("quickly reacting!")
}

func printMsg(msg string) {
	fmt.Println(msg)
}
