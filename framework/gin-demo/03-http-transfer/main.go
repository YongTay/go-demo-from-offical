package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	url := "https://raw.githubusercontent.com/YongTay/spotlight-in-chrome/main/package.json"
	r := gin.New()

	r.GET("/hello", func(c *gin.Context) {
		resp, err := http.Get(url)
		if nil != err {
			c.String(http.StatusInternalServerError, "")
			panic(err)
		}
		reader := resp.Body
		// 数据转发
		c.DataFromReader(200, resp.ContentLength, resp.Header.Get("Content-Type"), reader, map[string]string{ // 设置额外的响应头
			"user-kk": "transfer",
		})
	})
	r.Run(":8080")
}
