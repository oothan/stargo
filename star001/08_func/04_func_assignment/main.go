package main

import "fmt"

func main() {
	data := []string{"S快捷快递", "Kkk", "Data", "", "Data"}

	fmt.Printf("%q\n", trimSpace(data...))
	fmt.Printf("%q\n", data)
}

func trimSpace(data ...string) []string {
	var newData []string
	for _, d := range data {
		if d != "" {
			newData = append(newData, d)
		}
	}

	return newData
}
