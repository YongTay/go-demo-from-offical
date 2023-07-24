package main

import "fmt"

func main() {
	i := 21
	// case 多条件
	switch i {
	case 1, 11, 21:
		fmt.Println("i = 1")
	case 2:
		fmt.Println("i = 2")
	default:
		fmt.Println("default")
	}

	// 类似 if/else
	switch {
	case i < 10:
		fmt.Printf("i < 10\n")
	case i > 10:
		fmt.Println("i > 10")
	default:
		fmt.Println("i = 10")
	}

	var t interface{} = i

	// 类型判断
	switch t.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")
	}
}
