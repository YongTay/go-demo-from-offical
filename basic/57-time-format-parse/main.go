package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC1123)) // Fri, 21 Jul 2023 17:31:37 CST
	fmt.Println(t.Format(time.RFC3339)) // 2023-07-21T17:31:37+08:00
	fmt.Println(t.Format(time.RFC822))  // 21 Jul 23 17:31 CST
	fmt.Println(t.Format(time.RFC850))  // Friday, 21-Jul-23 17:31:37 CST

	// 日期的显示和unix一样
	fmt.Println("常用的日期格式", t.Format("2006-01-02 15:04:05"))

	t1, err := time.Parse(
		time.RFC3339,
		"2023-03-12T12:35:00+08:00",
	)
	if nil != err {
		panic(err)
	}
	fmt.Println(t1)

	t2, err2 := time.Parse(
		"2006-1-2 15:4:5", // 输入参数小于10的不需要补0
		// "2006-01-02 15:04:05", // 输入参数小于10的需要补0
		"2023-04-05 10:5:46",
	)
	if nil != err2 {
		panic(err2)
	}
	fmt.Println(t2)
}
