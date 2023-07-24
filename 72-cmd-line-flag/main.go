package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag 会读取 cmd 的参数 , usage 参数是帮助说明
	wordStr := flag.String("word", "foo", "a string")

	numStr := flag.Int("num", 10, "a number")

	boolStr := flag.Bool("fork", false, "a fork(bool false)")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word", *wordStr)
	fmt.Println("num", *numStr)
	fmt.Println("fork", *boolStr)
	fmt.Println("svar", svar)
}
