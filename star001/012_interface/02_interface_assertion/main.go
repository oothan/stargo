package main

import "fmt"

func main() {
	var message interface{} = 10

	s, ok := message.(string)
	if ok {
		fmt.Printf("%q %T \n", s, message)
	} else {
		fmt.Printf("value is not a string - %q %T \n", s, message)
	}
}
