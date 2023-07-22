package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"a", "d", "h", "c"}
	sort.Strings(strs)
	fmt.Println("sorted strs", strs)

	nums := []int{3, 5, 1, 2, 4}
	sort.Ints(nums)
	fmt.Println("sorted int", nums)

	sorted := sort.IntsAreSorted(nums)
	fmt.Println("sorted", sorted)
}
