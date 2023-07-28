package main

import (
	"fmt"
	"reflect"
)

type myType int

func main() {
	var a myType = 10
	aType := reflect.TypeOf(a)
	aValue := reflect.ValueOf(a)
	fmt.Println(aValue.CanFloat())

	// 获取指定值类型
	// fmt.Println(aValue.Int())

	switch aType.Kind() {
	case reflect.Int:
		fmt.Println(aValue.Int())
	case reflect.Float32, reflect.Float64:
		fmt.Println(aValue.Float())
	default:
	}

	// 返回声明的类型
	fmt.Println(aType.Name())
	// 返回基本类型
	fmt.Println(aType.Kind())

}
