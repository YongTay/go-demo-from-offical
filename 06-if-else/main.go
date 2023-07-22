package main

import "fmt"

func main() {
	if 7%2 == 1 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if n := 10; n < 10 {
		fmt.Println("n < 10")
	} else {
		fmt.Println("n >= 10")
	}
}
