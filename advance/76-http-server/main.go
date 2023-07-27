package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		handleGet(w, req)
	case http.MethodPost:
		handlePost(w, req)
	default:
		code := http.StatusBadRequest
		http.Error(w, http.ErrNotSupported.ErrorString, code)
	}
}

func handleGet(w http.ResponseWriter, req *http.Request) {
	// 解析url上的参数
	req.ParseForm()
	fmt.Println(req.Form)

	url := req.URL
	fmt.Println(req.Proto)
	fmt.Println(req.RequestURI)
	fmt.Println(url.RequestURI())
	fmt.Println(url.Query())
	fmt.Println(url.RawQuery)
	fmt.Fprint(w)
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	// 解析参数
	req.ParseForm()
	fmt.Println("form", req.Form)
	ctype := req.Header.Get("content-type")
	fmt.Println(ctype)
	if ctype == "application/json" {
		// 获取body参数
		sc := bufio.NewScanner(req.Body)
		var body []string
		for sc.Scan() {
			body = append(body, sc.Text())
		}
		var data map[string]any

		// 将字符串转成json对象
		json.Unmarshal([]byte(strings.Join(body, "")), &data)
		fmt.Println("json", data)
	}

	fmt.Fprint(w)
}
func main() {
	// get请求获取请求参数
	http.HandleFunc("/hello", hello)

	fmt.Println("listen on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
