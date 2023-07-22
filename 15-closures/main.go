package main

import "fmt"

func createSeq() func() int {
	n := 0
	return func() int {
		n += 1
		return n
	}
}

func main() {
	next := createSeq()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	
}
