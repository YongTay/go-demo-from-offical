package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	r := gin.New()

	r.GET("/port", func(c *gin.Context) {
		url := c.Request.Host
		c.String(200, "port: %s", url)
	})
	srv8080 := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	srv8081 := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	// 开启多个服务，服务必须开启在goroutine，否则会阻塞
	go func() {
		srv8080.ListenAndServe()
	}()
	go func() {
		srv8081.ListenAndServe()
	}()

	// 一直等待直到手动结束
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	<-sig
	fmt.Println("finish...")
}
