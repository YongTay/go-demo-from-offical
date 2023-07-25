package main

import (
	"embed"
	"fmt"
)

//go:embed folder/a.txt
var fileString string

//go:embed folder/a.txt
var fileByte []byte

// 将这个文件夹都导入进来
//
//go:embed folder/*
var folder embed.FS

func main() {
	// 直接获取导入文件的数据
	fmt.Println(fileString)
	fmt.Println(fileByte)

	a, err := folder.ReadFile("folder/a.txt")
	if nil != err {
		panic(err)
	}
	fmt.Println(a)

	// 读取指定的导入文件
	h1, err := folder.ReadFile("folder/hash1")
	if nil != err {
		panic(err)
	}
	fmt.Println(h1)
}
