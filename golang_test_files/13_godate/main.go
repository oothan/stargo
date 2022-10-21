package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	todayDate := currentTime.Format("2006-01-02")
	fmt.Println("YYYY-MM-DD : ", todayDate)

	var in int64
	in = 24

	if in == -1 {
		fmt.Println(in)
	} else if in <= 10 {
		fmt.Println(in)
	} else if in > 10 {
		fmt.Println(in)
	}
}
