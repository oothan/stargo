package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/peterhellberg/fixer"
)

func main() {
	f := flag.String("from", "USD", "")
	t := flag.String("to", "TRX", "")
	n := flag.Float64("n", 1, "")

	flag.Parse()

	from, to := fixer.Currency(*f), fixer.Currency(*t)

	resp, err := fixer.Latest(context.Background(),
		fixer.Base(from), fixer.Symbols(to),
	)

	if err == nil {
		fmt.Printf("%.2f %s equals %.2f %s\n", *n, from, resp.Rates[to]**n, to)
	}
}
