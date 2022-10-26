package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type money struct {
	Base     string  `json:"base"`
	Currency string  `json:"currency"`
	Amount   float32 `json:"amount"`
}

type info struct {
	Data []*money
}
type Range struct {
	Min     int64   `json:"min"`
	Max     int64   `json:"max"`
	Percent float64 `json:"percent"`
}

type RangeResponse struct {
	Data []*Range `json:"data"`
}

func main() {
	str := `{\"data\":[{\"min\": 1,\"max\": 1000,\"percent\": 0.01}]}`

	var i info

	if err := json.Unmarshal([]byte(str), &i); err != nil {
		fmt.Println("ugh: ", err)
	}

	m := &money{}
	for _, s := range i.Data {
		m = s
	}

	fmt.Println("info: ", i)
	fmt.Println("currency: ", m.Currency)

	tFormat := time.Now().Format("Mon Jan 02 15:04:05 -0700 2006")
	fmt.Println(tFormat)

}
