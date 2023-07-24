package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 获取所有的系统环境参数
	envs := os.Environ()
	for _, env := range envs {
		kv := strings.Split(env, "=")
		fmt.Println(kv[0], kv[1])
	}
	fmt.Println()
	v := os.Getenv("GOROOT")
	fmt.Println(v)
}
