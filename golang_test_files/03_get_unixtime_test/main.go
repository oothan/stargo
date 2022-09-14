package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2022-09-13"))

	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())
}
