package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	s1 := filepath.Join("a", "b", "c")
	fmt.Println(s1)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println(filepath.Dir(s1))
	fmt.Println(filepath.Base(s1))

	fmt.Println(filepath.IsAbs(s1))

	filename := "a.txt"

	fmt.Println(filepath.Ext(filename))

	// 获取相对路径
	s2, err := filepath.Rel("a/b", "a/b/c/d/path")
	if nil != err {
		panic(err)
	}
	fmt.Println(s2)

	s3, err := filepath.Rel("a/b", "a/d/c/d/path")
	if nil != err {
		panic(err)
	}
	fmt.Println(s3)
}
