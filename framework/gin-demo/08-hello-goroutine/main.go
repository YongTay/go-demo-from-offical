package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/hello-async", func(c *gin.Context) {
		// 副本拷贝
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			// 如果不使用副本，那么在context上下文中携带的数据，可能会影响到其它请求，因为context是复用的
			//c.Set("cCp", "this is from hello-async")
			// 使用副本
			log.Println("url: ", cCp.Request.RequestURI)
		}()
	})
	r.GET("/hello-sync", func(c *gin.Context) {
		v, ok := c.Get("cCp")
		log.Println("url: ", v, ok)
	})
	r.Run(":8080")
}
