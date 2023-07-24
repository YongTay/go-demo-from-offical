package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(name string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(name)
	if nil != err {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprint(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// 需要关注一下文件关闭的错误
	if nil != err {
		fmt.Fprintf(os.Stderr, "has error %v\n", err)
		os.Exit(1)
	}
}
