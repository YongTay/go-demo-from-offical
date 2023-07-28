package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int
}

func binding[T any](o *T, jsonStr string) {
	var data map[string]any
	err := json.Unmarshal([]byte(jsonStr), &data)
	if nil != err {
		panic(err)
	}
	// 等价于reflect.TypeOf(o).Elem()
	aType := reflect.TypeOf(*o)
	aValue := reflect.ValueOf(o)
	elem := aValue.Elem()
	n := elem.NumField()
	for i := 0; i < n; i++ {
		fv := elem.Field(i)
		f := aType.Field(i)
		var key = f.Name
		if k, ok := f.Tag.Lookup("json"); ok {
			key = k
		}
		val := data[key]
		// interface{} 不能强转成 int
		// 只能转成 float64 再强转成 int
		fmt.Printf("type: %t\n", val)
		if val != nil {
			switch fv.Kind() {
			case reflect.Int:
				fv.Set(reflect.ValueOf(int(val.(float64))))
			default:
				fv.Set(reflect.ValueOf(val))
			}
		}
	}

}

func main() {
	str := `{"name": "Tom", "Age": 20}`
	var p Person
	binding(&p, str)
	fmt.Println(p)
}
