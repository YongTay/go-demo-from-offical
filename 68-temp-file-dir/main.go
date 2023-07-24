package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if nil != e {
		panic(e)
	}
}
func main() {
	f, err := os.CreateTemp("", "sample")
	check(err)
	fmt.Println(f.Name())
	defer os.RemoveAll(f.Name())

	dname, err := os.MkdirTemp("", "sample")
	check(err)
	defer os.RemoveAll(dname)
	fmt.Println(dname)
}
