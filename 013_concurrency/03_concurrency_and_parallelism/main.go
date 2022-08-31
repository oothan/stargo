package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var waitG sync.WaitGroup
var cpuUsed = 1
var maxRandomNums = 20

func init() {
	maxCPU := runtime.NumCPU()
	cpuUsed = 4
	runtime.GOMAXPROCS(cpuUsed)

	fmt.Printf("Number of CPU (total=%d - Used=%d) \n", maxCPU, cpuUsed)
}

func main() {
	start := time.Now()
	ids := []string{"#", "!", "^", "*"}

	waitG.Add(4)
	for i := range ids {
		go numbers(ids[i])
	}
	waitG.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\nProgram took %s. \n", elapsed)
}

func numbers(id string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= maxRandomNums; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%s%d ", id, rand.Intn(20)+20)
	}
	waitG.Done()
}
