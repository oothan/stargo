package main

import (
	"fmt"
	"os"
	path2 "path"
)

func main() {
	dirPath, _ := os.Getwd()
	path := path2.Join(dirPath, "foo/klklkl", "l", "p")
	err := os.MkdirAll(path, 0777)
	fmt.Println(err)
}
