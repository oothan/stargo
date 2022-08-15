package main

import (
	"fmt"
	"math"
)

func main() {
	values := []string{"tea", "tae", "eta", "eat", "ate", "aet", "teal", "teatop", "teaparty", "tear"}

	for v := range values {
		hashKey := hashK(values[v])
		fmt.Printf(" %-10s (hashKey: %10d) %v \n", values[v], hashKey, []byte(values[v]))
	}
}

func hashK(s string) int {
	result := 0
	idx := 0
	for i := len(s) - 1; i > -1; i-- {
		result += int(math.Pow10(i)) * int(s[idx])
		idx++
	}
	return result
}
