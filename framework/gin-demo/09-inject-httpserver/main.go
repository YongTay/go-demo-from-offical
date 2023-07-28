package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome",
		})
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	srv.ListenAndServe()
}
