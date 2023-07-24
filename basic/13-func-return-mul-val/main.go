package main

import "fmt"

func mm() (int, int) {
	return 5, 19
}

// 可以结构其它函数返回的多个参数
func mulParam(i, j int) {
	fmt.Println(i, j)
}

func main() {
	a, b := mm()

	fmt.Println(a, b)
	mulParam(mm())
}
