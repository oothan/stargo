package main

import "fmt"

var i = 0

func main() {
	fmt.Println("i(m1) = ", i)
	testDefer()
	fmt.Println("i(m2) = ", i)
}

func testDefer() {
	fmt.Println("i(t1) = ", i)

	defer closeFiles()
	defer closeDBConnections()

	fmt.Println("i(t2) = ", i)
	doSomething()
	fmt.Println("i(t3) = ", i)
}

func closeFiles() {
	fmt.Println("i(f1) = ", i)
	i = 1
}

func closeDBConnections() {
	fmt.Println("i(f2) = ", i)
	i += 2
}

func doSomething() {
	fmt.Println("i(f3) = ", i)
	i = 3
}
