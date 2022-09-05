package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan string)

	start := time.Now()

	go message1(c)
	go message2(c)
	go message3(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	close(c)
	elapsed := time.Since(start)
	fmt.Printf("\ntime elapsed: %s\n", elapsed)
}

func message1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "msg1, delay 3 sec"
}

func message2(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "msg2, delay 2 sec"
}

func message3(c chan string) {
	time.Sleep(4 * time.Second)
	c <- "msg3, delay 4 sec"
}
