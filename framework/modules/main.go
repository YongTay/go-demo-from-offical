package main

import (
	"fmt"
	"os"
	"strings"

	"framework/modules/utils"
)

func main() {
	res := utils.Add(1, 2, 3)
	fmt.Println(res)

	// dirs, err := os.ReadDir("/Applications")
	// if nil != err {
	// 	panic(err)
	// }
	// for _, v := range dirs {
	// 	fmt.Println(v.Name())
	// }
	// iterm, err := exec.LookPath("ls")
	// if nil != err {
	// 	panic(err)
	// }
	// err = syscall.Exec(iterm, []string{"ls"}, os.Environ())
	// if nil != err {
	// 	panic(err)
	// }

	pathEnv := os.Getenv("PATH")
	path := strings.Split(pathEnv, ":")
	for _, v := range path {
		fmt.Println(v)
		files, err := os.ReadDir(v)
		if nil != err {
			panic(err)
		}
		for _, f := range files {
			fmt.Println(f.Name())
		}
	}
}
