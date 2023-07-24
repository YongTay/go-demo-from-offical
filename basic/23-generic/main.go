package main

import "fmt"

// K: 输入的参数 comparable: 限定的参数
func mapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type element[T any] struct {
	val  T
	next *element[T]
}

type List[T any] struct {
	// 链表头，链表尾
	header, tail *element[T]
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.header = &element[T]{val: v}
		list.tail = list.header
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list List[T]) GetAll() []T {
	var r []T
	for e := list.header; e != nil; e = e.next {
		r = append(r, e.val)
	}
	return r
}

func main() {
	m := map[string]int{"A": 10, "B": 11, "C": 23}
	fmt.Println(mapKeys(m))
	fmt.Println(mapKeys[string, int](m))

	list := List[int]{}
	list.Push(10)
	list.Push(23)
	list.Push(14)
	all := list.GetAll()
	fmt.Println(all)
}
