package main

import "fmt"

func main() {
	factor := 3

	mult := func(i, j int) int {
		return i * j * factor
	}

	t := mult(3, 4)
	fmt.Println(t)
	fmt.Println(mult(4, 5))

	hi := sayGreeting("ESP")
	fmt.Println(hi())

	fmt.Println(sayGreeting("ENG")())

	fmt.Println()
	fmt.Printf("%T\n", hi())
	fmt.Printf("%T\n", hi)

}

func sayGreeting(lang string) func() string {
	hi := func() string {
		return "hi"
	}

	if lang == "ESP" {
		hi = func() string {
			return "hola"
		}
	} else if lang == "GER" {
		hi = func() string {
			return "hello"
		}
	}

	return hi
}
