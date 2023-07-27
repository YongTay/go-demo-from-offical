package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome",
		})
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); nil != err && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 优雅的关闭
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	// 等待退出指令
	<-quit
	fmt.Println("shutdown...")
	// 退出之后的处理
	// 超时5秒后自动退出
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	if err := srv.Shutdown(ctx); nil != err {
		log.Fatalf("shutdown Error: %s\n", err)
	}
	log.Println("shutdown server")

	defer func() {
		fmt.Println("正常退出")
		cancel()
	}()

}
