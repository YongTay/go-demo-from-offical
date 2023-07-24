package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	// 0 - 启动文件名 1: 输入参数
	fmt.Println(args)
	fmt.Println(args[1:])
}
