package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func Server() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
