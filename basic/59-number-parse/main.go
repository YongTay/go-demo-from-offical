package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.23", 64)
	fmt.Println(f)

	// base 0 表示自动根据输入的参数进行自动判断，比如0x,0o， 0b...
	// 否则将按照指定的输入进行解析
	n1, _ := strconv.ParseInt("14", 0, 64)
	fmt.Println(n1)
	n2, _ := strconv.ParseInt("0x14", 0, 64)
	fmt.Println(n2)
	n3, _ := strconv.ParseInt("14", 16, 64)
	fmt.Println(n3)

	// 10进制转换简化
	n4, _ := strconv.Atoi("123")
	fmt.Println(n4)
}
