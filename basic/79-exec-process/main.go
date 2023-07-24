package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	bin, err := exec.LookPath("ls")
	if nil != err {
		panic(err)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	// 采用系统调用
	execErr := syscall.Exec(bin, args, env)
	if nil != execErr {
		panic(err)
	}
}
