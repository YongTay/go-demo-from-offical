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
	data, err := os.ReadFile("./a.txt")
	check(err)
	fmt.Println(data)

	f, err := os.Open("./a.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Println("读到的数据长度", n1)
	fmt.Println(string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Println("n2", n2, "o2", o2)
	fmt.Println(string(b2[:n2]))

	// 重置观察点
	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Println(string(b4))
}
