package main

import (
	"fmt"
	"strings"
)

func main() {
	sourceString := "Error 1062: Duplicate entry 'tester2' for key 'domain_groups.name'"
	//sourceString = strings.ReplaceAll(sourceString, " ", "")
	fmt.Println(strings.HasPrefix(sourceString, "Error 1062"))
	fmt.Println(strings.HasSuffix(sourceString, "'domain_groups.name'"))

	sourceString1 := "Error 1062: Duplicate entry 'sai.com' for key 'domain_groups.url'"
	//sourceString = strings.ReplaceAll(sourceString, " ", "")
	fmt.Println(strings.HasPrefix(sourceString1, "Error 1062"))
	fmt.Println(strings.HasSuffix(sourceString1, "'domain_groups.url'"))
}
