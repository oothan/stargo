package main

import (
	"fmt"
	"math"
	"math/rand"
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
	//s := "1~5"
	//spl := s	trings.Split(s, "~")
	//minn, _ := strconv.ParseInt(spl[0], 10, 64)
	//maxx, _ := strconv.ParseInt(spl[1], 10, 64)
	//fmt.Println(minn, " ", maxx)

	//rng := strings.Split(s, "~")
	//min, _ := strconv.ParseInt(rng[0], 10, 64)
	//max, _ := strconv.ParseInt(rng[1], 10, 64)
	//rand.Seed(time.Now().UnixNano())
	//per := rand.Int63n(max-min) + min
	//
	//fmt.Println(per)
	//fmt.Println(float64(per) * 0.01)

	t := 0.01
	tt := t * 100
	tt1 := tt + 0.3
	tt2 := tt - 0.3

	ttt1 := tt1 / 100
	ttt2 := tt2 / 100

	fmt.Println(tt1, " ", tt2)
	fmt.Println(ttt1, " ", ttt2)
	fmt.Println(roundFloat(ttt1, 4), " ", roundFloat(ttt2, 4))

	rand.Seed(time.Now().UnixNano())
	r := ttt1 + rand.Float64()*(ttt2-ttt1)
	fmt.Println(r)
	fmt.Println(roundFloat(r, 4))

	g := float64(0.7)
	fmt.Println(int64(g))
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
