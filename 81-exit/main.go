package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os.exit 中断程序，defer将无法执行
	defer fmt.Println("!")

	// 系统异常级别中断退出
	os.Exit(3)
}
