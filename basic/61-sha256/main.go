package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "this is a string"
	h := sha256.New()

	// 计算hash
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
