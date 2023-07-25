package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if nil != e {
		panic(e)
	}
}
func main() {
	err := os.Mkdir("tmp", 0755)
	check(err)
	defer os.RemoveAll("tmp")

	createEmptyFile := func(name string) {
		err := os.WriteFile(name, []byte(""), 0644)
		check(err)
	}

	createEmptyFile("tmp/file1")

	// 创建多层文件
	err = os.MkdirAll("tmp/sub", 0755)
	check(err)

	createEmptyFile("tmp/sub/file2")
	createEmptyFile("tmp/sub/file3")
	createEmptyFile("tmp/sub/file4")

	dir, err := os.ReadDir("tmp/sub")
	check(err)
	for _, d := range dir {
		fmt.Println(d.Name(), d.IsDir())
	}

	// 遍历文件夹及其子文件
	filepath.Walk("tmp", func(path string, info fs.FileInfo, err error) error {
		if nil != err {
			return err
		}
		fmt.Println(path, info.IsDir())
		return nil
	})
}
