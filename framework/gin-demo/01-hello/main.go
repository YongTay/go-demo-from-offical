package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 使用中间件，会输出日志信息
	// r := gin.Default()
	// 不使用默认中间件，不会有日志信息
	r := gin.New()
	r.GET("/hello", func(c *gin.Context) {
		fmt.Println(c.Query("name"))
		fmt.Println(c.Query("age"))
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome",
		})
	})
	r.Run(":8080")
}
