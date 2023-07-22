package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "中国"

	fmt.Println("len: ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()

	fmt.Println("Rune Len: ", utf8.RuneCountInString(s))

	for _, val := range s {
		// fmt.Printf("%#U", val)

		// rune 可以直接输出内容
		fmt.Printf("%c", val)
	}
	fmt.Println()

	for i, w := 0, 0; i < len(s); i += w {
		val, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U", val)
		w = width
	}
	fmt.Println()

}
