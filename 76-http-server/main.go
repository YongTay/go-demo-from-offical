package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "%v, %v", k, v)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	fmt.Println("listen on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
