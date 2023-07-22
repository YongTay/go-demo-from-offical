package main

import (
	"fmt"
	"regexp"
)

func main() {

	// direct to match
	match, _ := regexp.MatchString("^p[a-z]+ch$", "peach")
	fmt.Println(match)

	// new a complier
	// 需要自行处理error
	r, _ := regexp.Compile("p[a-z]+h")
	// 方法内部处理error
	// mr := regexp.MustCompile("")

	fmt.Println(r.MatchString("peaaaaaach"))

	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindAllString("peach punch", -1))

	fmt.Println(r.FindStringIndex("punch peach"))
	fmt.Println(r.FindAllStringIndex("punch peach", -1))

	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindAllStringSubmatch("peach punch", -1))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch", -1))

	fmt.Println(r.Match([]byte("peach")))
}
