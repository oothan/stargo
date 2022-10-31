package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Add(time.Hour * 6)
	fmt.Println(t)

}
