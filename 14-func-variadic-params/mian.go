package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println(nums)

	sum := 0

	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	sum0 := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(sum0)

	arr := []int{10, 11, 12}
	sum1 := sum(arr...)
	fmt.Println(sum1)
}
