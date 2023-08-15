package main

import (
	"fmt"
	"unsafe"
)

type tPoint struct {
	data int
	name string
	age  int
}

func change(a *int) int {
	*a = 20
	return *a
}
func main() {
	var a int
	fmt.Println(a)

	change(&a)
	fmt.Println(a)

	var ptr *int = nil
	fmt.Println(ptr) // nil
	// fmt.Println(*ptr) // nil 不允许取值

	ptr = new(int)
	fmt.Println(ptr)  // 0x140000140e8
	fmt.Println(*ptr) // 初始值与new 的类型一样 0

	*ptr = 100
	fmt.Println(*ptr) // 100
	// uintptr 允许指针相加，但是无法直接取值，必须转成具体类型才能获取到实际的值
	// unsafe.Pointer 不允许进行指针相加，但是可以转成uintptr，然后让uintptr进行指针相加等操作，同时也可转成具体类型，读取实际的值
	// uintptr <=> unsafe.Pointer <=> *T
	// unsafe.Pointer 就是一个中转站
	s := make([]int, 3, 20)
	// make 返回的结构体
	// { typeArray []int, Len int , Cap int }
	// []int 再内存中的内存地址也是一个int
	// 下面的uintptr(8)指的是一个int大小
	// uintptr(16) 是指2个int大小，这里是指需要跨越2个int，来访问到Cap
	fmt.Printf("%v\n", &s)
	pointer := unsafe.Pointer(&s)        // struct 强转成 unsafe.Pointer
	Len := uintptr(pointer) + uintptr(8) // uintptr 指针相加
	upval := unsafe.Pointer(Len)         // 转成具体类型
	val := *(*int)(upval)                // 强转访问具体值
	// 合并写法 *(*int) 的说明
	Cap := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(pointer)
	fmt.Println(val)
	fmt.Println(Cap)

	//
	tp := tPoint{10, "John", 20}
	data := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&tp)) /* + uintptr(0) */))
	name := *(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&tp)) + uintptr(8)))
	age := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&tp)) + uintptr(24)))
	fmt.Println("data", data)
	fmt.Println("name", name)
	fmt.Println("age", age)
	fmt.Println("修改name名称")
	// 根据结构题类型，name前面有一个int类型
	// 如果你足够自自信，那么直接uintptr(8)即可，这就表示一个int的大小
	// 当然你也可以使用unsafe.Sizeof函数来计算int的大小，当然其它类型也类似
	ptrName := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&tp)) + unsafe.Sizeof(int(0))))
	ptrAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&tp)) + unsafe.Offsetof(tp.age))) // 通过字段名直接获取偏移量
	//ptrAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&tp)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))
	fmt.Printf("sizeof int: %d, size of string: %d, offset age: %d\n", unsafe.Sizeof(0), unsafe.Sizeof(""), unsafe.Offsetof(tp.age))
	fmt.Println("修改前：", tp)
	*ptrName = "Tom"
	*ptrAge = 23
	fmt.Println("修改后：", tp)
	fmt.Println(unsafe.Alignof(tp))
}
