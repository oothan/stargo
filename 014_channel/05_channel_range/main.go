package main

import "fmt"

var wordSet1 = []string{"small", "medium", "large"}
var wordSet2 = []string{"red", "green", "blue", "yellow"}

func main() {
	c := make(chan string)

	go queue1(c)
	go queue2(c)

	for val := range c {
		fmt.Println(val)
	}
}

func queue1(c chan string) {
	for _, w := range wordSet1 {
		c <- w
	}
}

func queue2(c chan string) {
	for _, w := range wordSet2 {
		c <- w
	}
	close(c)
}
