package main

import (
	"go-accountbook/registry"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// historyCtr := registry.InjectionHistoryCtr()

	historyCtl := registry.InjectionHistory()

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Use(cors.Default())
	server.GET("/history/:id", historyCtl.GetHistory)
	server.GET("/history", historyCtl.GetHistories)

	server.Run()
}
