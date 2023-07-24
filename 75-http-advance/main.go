package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http advance")
	client := &http.Client{}
	resp, err := client.Get("https://www.baidu.com")
	if nil != err {
		panic(err)
	}
	defer resp.Body.Close()

	sc := bufio.NewScanner(resp.Body)
	for sc.Scan() {
		// 输出为string
		// fmt.Println(sc.Text())
		// 查找指定字符
		n := bytes.Index(sc.Bytes(), []byte("l"))
		fmt.Println(n)
	}
}
