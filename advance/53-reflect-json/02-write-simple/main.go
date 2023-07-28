package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	n := &a
	fmt.Println("before change", a)
	aValue := reflect.ValueOf(n)
	aValue.Elem().Set(reflect.ValueOf(20))
	fmt.Println("after changed", a)

}
