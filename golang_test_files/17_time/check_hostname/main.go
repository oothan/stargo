package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	input := "https://www.gosamples.dev/abc/def"

	url, err := url.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	hostname := strings.TrimPrefix(url.Hostname(), "www.")

	fmt.Println(hostname)

	str2 := "@@This is the tutorial of Golang$$"
	fmt.Println(strings.ReplaceAll(str2, " ", ""))
}
