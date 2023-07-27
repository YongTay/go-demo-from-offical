package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func muAuthorized() func(*gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("authorized...")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// 创建路由组
	authorized := r.Group("/admin", muAuthorized())
	// curl http://localhost:8080/admin/ping
	authorized.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	abc := r.Group("/a").Group("/b").Group("c")
	// curl http://localhost:8080/a/b/c/ping
	abc.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
