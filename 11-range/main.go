package main

import "fmt"

func main() {
	line := "hello world"
	// index, value
	for k, v := range line {
		fmt.Println(k, v)
	}

	m := map[string]string{"a": "h", "b": "ee"}
	// key value
	for k, v := range m {
		fmt.Println(k, v)
	}
}
