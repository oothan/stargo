package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	timeNano := time.Now().UnixNano()
	timeUnix := time.Now().Unix()
	timeMili := time.Now().UnixMilli()
	timeMicro := time.Now().UnixMicro()
	fmt.Println(timeNano)
	fmt.Println(timeUnix)
	fmt.Println(timeMili)
	fmt.Println(timeMicro)

	fmt.Println(rand.Uint32())

	fmt.Println(strconv.FormatInt(time.Now().UnixNano()+int64(rand.Uint32()), 10))
	dt := time.Now()
	sd := dt.Day()
	sy := dt.Year()
	sm := int64(dt.Month())
	fmt.Println(fmt.Sprintf("%v %v %v", sd, sm, sy))

	firstDate := time.Date(2022, 4, 29, 0, 0, 0, 0, time.UTC)
	secondDate := time.Date(2022, 4, 30, 0, 0, 0, 0, time.UTC)
	difference := firstDate.Sub(secondDate)

	fmt.Println(int64(difference.Hours() / 24))

}
