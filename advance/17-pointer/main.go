package main

import "fmt"

func change(a *int) int {
	*a = 20
	return *a
}
func main() {
	var a int
	fmt.Println(a)

	change(&a)

	fmt.Println(a)
}
