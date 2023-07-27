package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
)

func check(e error) {
	if nil != e {
		panic(e)
	}
}
func main() {
	url, _ := url.Parse("http://127.0.0.1:7890")
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(url),
		},
	}
	resp, err := client.Get("https://www.google.com")

	check(err)
	defer resp.Body.Close()

	sc := bufio.NewScanner(resp.Body)

	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
