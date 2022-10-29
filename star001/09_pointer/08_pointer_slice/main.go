package main

import "fmt"

func main() {
	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Printf("days=%s\n", days)
	fmt.Printf("len(days)=%d\n", len(days))
	fmt.Printf("cap(days)=%d\n", cap(days))

	weekday := days[3:6]
	fmt.Printf("days=%s\n", weekday)
	fmt.Printf("len(days)=%d\n", len(weekday))
	fmt.Printf("cap(days)=%d\n", cap(weekday))

	fmt.Printf("&days[3]=%x\n", &days[3])
	fmt.Printf("&weekday[0]=%x\n", &weekday[0])

	fmt.Println()
	fmt.Printf("&days[5]=%x\n", &days[5])
	fmt.Printf("&weekday[2]=%x\n", &weekday[2])
}
