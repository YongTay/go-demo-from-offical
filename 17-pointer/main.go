package main

import "fmt"

func zeroval(n int) {
	n = 0
}

func zeroptr(n *int) {
	*n = 0
}

func main() {
	i := 1
	// 值传递
	zeroval(i)
	fmt.Println(i)

	// 内存地址传递
	zeroptr(&i)
	fmt.Println(i)
}
