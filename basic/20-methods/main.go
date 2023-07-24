package main

import "fmt"

type rect struct {
	width, height int
}

// 引用传递
func (r *rect) area() int {
	return r.width * r.height
}

// 没有使用指针类型，那么外部的改动将不会生效
// 值传递
func (r rect) perim() int {
	return 2*r.height + 2*r.height
}
func main() {
	r := rect{5, 6}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	r2 := &r
	r2.width = 6
	fmt.Println(r2.area())
	fmt.Println(r.perim())
}
