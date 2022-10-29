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
}
