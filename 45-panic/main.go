package main

import "fmt"

func main() {
	// abort
	panic("it has some panic")

	fmt.Println("will arrive?") // no
}
