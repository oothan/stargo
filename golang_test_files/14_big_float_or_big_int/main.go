package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	//t1 := time.UnixMilli(1664457801000)
	//t := t1.Format("2006-01-02")
	//fmt.Println(t)
	//
	//s := "12450000"
	//f, _ := strconv.ParseInt(s, 10, 64)
	//fmt.Println(f)
	//fl := float64(f) / 1000000
	//fmt.Println(fl)

	//min := 1
	//max := 5
	//// set seed
	//rand.Seed(time.Now().UnixNano())
	//// generate random number and print on console
	//per := rand.Intn(max-min) + min

	//fmt.Println(per)
	//fmt.Println(float64(per) * 0.01)
	//
	s := "1~5"
	//spl := s	trings.Split(s, "~")
	//minn, _ := strconv.ParseInt(spl[0], 10, 64)
	//maxx, _ := strconv.ParseInt(spl[1], 10, 64)
	//fmt.Println(minn, " ", maxx)

	rng := strings.Split(s, "~")
	min, _ := strconv.ParseInt(rng[0], 10, 64)
	max, _ := strconv.ParseInt(rng[1], 10, 64)
	rand.Seed(time.Now().UnixNano())
	per := rand.Int63n(max-min) + min

	fmt.Println(per)
	fmt.Println(float64(per) * 0.01)
}
