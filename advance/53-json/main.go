package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
}

func check(err error) {
	if nil != err {
		panic(err)
	}
}

func main() {
	f, err := os.OpenFile("./test.txt", os.O_RDONLY, 0755)
	check(err)
	r := bufio.NewReader(f)

	dec := json.NewDecoder(r)
	var p Person
	err = dec.Decode(&p)
	if nil != err {
		panic(err)
	}
	
	fmt.Println(p)
}
