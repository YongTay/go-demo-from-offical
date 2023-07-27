package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

type PersonJson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/info", func(c *gin.Context) {
		var p Person
		// 底层使用放射进行绑定
		if err := c.ShouldBindQuery(&p); nil != err {
			fmt.Println(err)
		}
		fmt.Println(p)
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.POST("/info", func(c *gin.Context) {
		var p PersonJson
		// 底层使用json进行绑定
		if err := c.ShouldBindJSON(&p); err != nil {
			fmt.Println(err)
		}
		fmt.Println(p)
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.Run(":8080")
}
