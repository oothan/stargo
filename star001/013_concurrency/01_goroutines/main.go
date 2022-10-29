package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//numbers()
	//alphabets()

	go numbers()
	go alphabets()

	time.Sleep(3200 * time.Millisecond)
	fmt.Println("\nmain terminated")
}

func numbers() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 15; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", rand.Intn(20)+20)
	}
}

func alphabets() {
	for i := 'C'; i <= 'G'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
