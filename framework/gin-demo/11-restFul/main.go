package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 匹配/hello/xxx
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "welcome", name)
	})
	r.GET("/info/:sex/:from", func(c *gin.Context) {
		sex := c.Param("sex")
		from := c.Param("from")
		c.String(200, "sex: %s, from: %s\n", sex, from)
	})
	r.Run(":8080")
}
