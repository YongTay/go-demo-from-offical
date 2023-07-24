package main

import "fmt"

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntSum(a ...int) int {
	sum := 0
	for _, n := range a {
		sum += n
	}
	return sum
}

func main() {
	fmt.Println(IntMin(10, 3))
}
