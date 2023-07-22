package main

import "fmt"

func main() {

	// 对panic进行收尾处理
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover handler!", r)
		}
	}()

	panic("an error")

	fmt.Println("will arrive ?") // no
}
