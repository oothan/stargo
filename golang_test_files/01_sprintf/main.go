package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, age = "Kiki", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)

	io.WriteString(os.Stdout, s)
}
