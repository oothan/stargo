package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitG sync.WaitGroup

func main() {
	//numbers()
	//alphabets()

	waitG.Add(2)
	go numbers()
	go alphabets()
	waitG.Wait()

	//time.Sleep(3200 * time.Millisecond)
	fmt.Println("\nmain terminated")
}

func numbers() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 15; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", rand.Intn(20)+20)
	}
	waitG.Done()
}

func alphabets() {
	for i := 'C'; i <= 'G'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	waitG.Done()
}
