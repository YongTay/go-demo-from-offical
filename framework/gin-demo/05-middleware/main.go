package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 自定义中间件
func myMiddle(name string) func(*gin.Context) {
	// 启动时会注册调用
	fmt.Println("myMiddle...")
	// TODO...
	return func(c *gin.Context) {
		// 每次都会调用
		fmt.Println(name, "myMiddle next...")
		// 放行，
		c.Next()

		fmt.Println(name, "myMiddle next after...")
	}
}

// 自定义中间件
func myMiddle2(name string) func(*gin.Context) {
	// 启动时会注册调用
	fmt.Println("myMiddle...")
	// TODO...
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		// 每次都会调用
		fmt.Println(name, "myMiddle2 next...", uri)

		// 放行，
		c.Next()
		// 其它中间件执行完才到这
		fmt.Println(name, "myMiddle2 next after...", uri)
	}
}

func main() {
	r := gin.New()
	// 全局中间件
	r.Use(gin.Logger())
	r.Use(myMiddle("全局"))
	r.Use(myMiddle2("全局"))
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome",
		})
	})
	// 注册局部中间件
	// 局部中间件会最后执行
	r.GET("/ping", myMiddle2("指定"), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
