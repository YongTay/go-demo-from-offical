package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{3, 5}
	fmt.Printf("struct1 %v\n", p)  // struct1 {3 5}
	fmt.Printf("struct1 %+v\n", p) // struct1 {x:3 y:5}
	fmt.Printf("struct1 %#v\n", p) // struct1 main.point{x:3, y:5}

	fmt.Printf("type %T\n", p)
	fmt.Printf("bool %t\n", true)
	fmt.Printf("int %d\n", 123)
	fmt.Printf("bin %b\n", 14)
	fmt.Printf("char %c\n", 97) // 97 => a ascii
	fmt.Printf("hex %x\n", 456)
	fmt.Printf("float %f\n", 12.45)
	fmt.Printf("float %e\n", 123400000.990)
	fmt.Printf("float %E\n", 123400000.990)
	fmt.Printf("str %s\n", "string")
	fmt.Printf("str %q\n", "string")
	fmt.Printf("str %x\n", "string")
	fmt.Printf("point %p\n", &p)
	fmt.Printf("width |%6d|%6d|\n", 12, 345)
	fmt.Printf("width |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("width |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width |%-6s|%-6s|\n", "foo", "b")

	// 返回结果
	s := fmt.Sprintf("this is a string: %s", "hallo word")
	fmt.Println(s)

	// 直接输出到流
	fmt.Fprintf(os.Stderr, "occur an error %s\n", "panic")
}
