package main

import "fmt"

func main() {
	var m = make(map[string]int)
	m["a"] = 1
	fmt.Println(m)

	var m1 = map[string]int{"a": 1, "b": 2, "c": 3}
	delete(m1, "a")
	fmt.Println(m1)

	// val, exist := m[key]
	_, b := m["a"]
	if b {
		fmt.Println(b, m["a"])
	} else {
		fmt.Println(b)
	}

	v := m1["c"]
	fmt.Println(v)

	fmt.Println("len(m1): ", len(m1))
}
