package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 接口实现
func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// 接口实现
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 面向接口操作
func measure(g geometry) {
	fmt.Println("data: ", g)
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}

func main() {
	r := rect{5, 6}
	c := circle{5}

	measure(r)
	measure(c)
}
