package main

import "fmt"

func main() {
	const s = 1
	// const 变量不允许再次赋值
	// s = ""

	var a int = s
	a = 2
	fmt.Println(a)

	fmt.Println(s)
}
