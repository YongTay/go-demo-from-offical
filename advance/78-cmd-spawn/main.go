package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 可以拿到执行结果的返回参数
	lsCmd := exec.Command("man", "ls")
	lsRes, err := lsCmd.Output()
	if nil != err {
		panic(err)
	}
	fmt.Println(string(lsRes))

	// 系统调用，将不会有返回参数
	bin, err := exec.LookPath("man")
	if nil != err {
		panic(err)
	}
	args := []string{"man", "ls"}
	err = syscall.Exec(bin, args, os.Environ())
	if nil != err {
		panic(err)
	}

}
