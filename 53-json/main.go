package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	MustMarshal := func(v any) []byte {
		data, err := json.Marshal(v)
		if nil != err {
			panic(err)
		}
		return data
	}

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB := MustMarshal(10)
	fmt.Println(string(intB))

	fltB := MustMarshal(1.23)
	fmt.Println(string(fltB))

	strB := MustMarshal("string")
	fmt.Println(string(strB))

	slcB := MustMarshal([]string{"apple", "peach", "pear"})
	fmt.Println(string(slcB))

	mapB := MustMarshal(map[string]any{"name": "Jane", "age": 23})
	fmt.Println(string(mapB))

	res1 := MustMarshal(response1{
		Page:   200,
		Fruits: []string{"apple", "orange", "purple"},
	})
	fmt.Println(string(res1))

	res2 := MustMarshal(&response2{
		Page:   500,
		Fruits: []string{"apple", "orange", "purple"},
	})
	fmt.Println(string(res2))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var data map[string]any
	if err := json.Unmarshal(byt, &data); nil != err {
		panic(err)
	}
	// 类型断言
	var f64 float64 = data["num"].(float64)
	fmt.Println(f64)

	// 设置输出流
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(map[string]int{"apple": 5, "lettuce": 7})

}
