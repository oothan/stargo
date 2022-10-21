package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	bal := 1.23
	decimals := 18
	fbal := new(big.Float)
	fbal.SetString(fmt.Sprintf("%f", bal))
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(decimals)))
	fmt.Println(value, " ", fmt.Sprintf("%f", bal))
}
