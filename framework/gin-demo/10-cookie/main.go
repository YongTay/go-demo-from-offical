package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 关闭控制台日志颜色
	gin.DisableConsoleColor()
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {

		ginCookie, err := c.Cookie("gin_cookie")
		if nil != err {
			c.SetCookie("gin_cookie", "power by gin", 360, "/", "localhost", false, true)
		}
		fmt.Println("cookie", ginCookie)
	})
	r.Run(":8080")
}
