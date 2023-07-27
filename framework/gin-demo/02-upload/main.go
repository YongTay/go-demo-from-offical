package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.MaxMultipartMemory = 8 * 1024 * 1024
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)
		// 将上传的文件保存到本地
		// c.SaveUploadedFile(file, "./"+file.Filename)
		c.String(200, fmt.Sprintf("upload %s", file.Filename))
	})

	r.Run(":8080")
}
