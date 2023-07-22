package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base's num is %d", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := &container{
		base: base{
			num: 20,
		},
		str: "some name",
	}

	fmt.Printf("co{num: %d, str: %s}\n", co.num, co.str)

	fmt.Printf("num is %d\n", co.base.num)

	fmt.Println("describe ", co.describe())

	// 后定义的接口类型也可以去接受已经存在的结构，实现多态
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describe ", d.describe())

}
