package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s, s == nil, len(s), cap(s))
	s = make([]string, 3)
	s[0] = "1"
	s[1] = "1"
	s[2] = "1"
	fmt.Println(s, s == nil, len(s), cap(s))

	var s1 = make([]string, len(s))
	copy(s1, s)
	fmt.Println(s1, s1 == nil, len(s1), cap(s1))
	fmt.Println(s[:])
}
