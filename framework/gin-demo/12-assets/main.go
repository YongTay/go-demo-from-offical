package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Static("/js", "./")
	r.Static("/img", "./")
	r.Run(":8080")
}
