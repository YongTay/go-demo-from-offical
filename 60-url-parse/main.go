package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if nil != err {
		panic(err)
	}
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.Query())
	fmt.Println(u.Fragment)
}
