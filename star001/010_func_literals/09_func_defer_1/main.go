package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("LOC 1", time.Now())

	//showMsg()
	defer showMsg()

	fmt.Println("LOC 2", time.Now())
	time.Sleep(time.Second * 1)
	fmt.Println("LOC 3", time.Now())
}

func showMsg() {
	fmt.Println("\nshowMsg", time.Now())
}
