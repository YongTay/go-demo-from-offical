package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if nil != e {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("./a.txt", d1, 0644)
	check(err)

	f, err := os.Create("./b.txt")
	check(err)
	defer f.Close()

	n1, err := f.Write([]byte("hello\njava\n"))
	check(err)
	fmt.Println("n1", n1)

	// 追加模式 append
	n2, err := f.WriteString("this is go demo")
	check(err)
	fmt.Println("n2", n2)

	w := bufio.NewWriter(f)
	// 追加模式 append
	n3, err := w.WriteString("how to do ")
	check(err)
	fmt.Println("n3", n3)

	// 刷新一下缓存，让数据立即写入
	w.Flush()
}
