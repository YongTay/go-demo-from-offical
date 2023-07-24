package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func check(e error) {
	if nil != e {
		panic(e)
	}
}

func main() {
	resp, err := http.Get("https://www.baidu.com")
	check(err)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
