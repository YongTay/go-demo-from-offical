package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 监听输入
func main() {
	fmt.Print("input:")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(strings.ToUpper(text))
		fmt.Print("input:")
	}

	if err := scanner.Err(); nil != err {
		panic(err)
	}
}
