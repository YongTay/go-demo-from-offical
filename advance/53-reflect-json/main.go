package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	aType := reflect.TypeOf(a)
	fmt.Println(aType.Name())
}
