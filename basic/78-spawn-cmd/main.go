package main

import (
	"fmt"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateRes, err := dateCmd.Output()
	if nil != err {
		panic(err)
	}
	fmt.Println(string(dateRes))
}
